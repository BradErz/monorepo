#!/bin/bash
# usage
# ./cache.sh restore $GOPATH/pkg/mod/cache/download go.sum
# ./cache save $GOPATH/pkg/mod/cache/download go.sum
# this is a generic script which when invoked with save will compress a specified folder and upload it to a bucket if
# it doesnt already exist.
# the restore operation will check if a file with the same hash exists on the bucket and if so downloads and extracts
# that file to the same local folder.

cache_version="v3"

bucket_name="default-cache-bucket"
if [[ -n ${BUCKET_NAME} ]]; then bucket_name=${BUCKET_NAME}; fi
cache_bucket_url="gs://${bucket_name}/cache/${PROJECT_NAME}"

task=$1
folders=$2
lock_file=$3


if [[ -z "$task" ]]; then echo "task must be specified" && exit 1; fi
if [[ -z "$folders" ]]; then echo "folders must be specified" && exit 1; fi
if [[ ! -f ${lock_file} ]]; then echo "${lock_file}  must exist" && exit 1; fi

lock_checksum=$(md5sum "${lock_file}"  | cut -d ' ' -f 1)

cache_file_name="${lock_checksum}-${cache_version}.tar.gz"
cache_file_url="$cache_bucket_url/$cache_file_name"

gsutil -q stat "${cache_file_url}"
cache_file_exists=$?


case "$task" in
    restore)
        if [[ "$cache_file_exists" == "0" ]]; then
            echo "Found cache $cache_file_name at $cache_bucket_url"

            echo "Downloading cache ..."
            time gsutil cp "${cache_file_url}" - | tar xzf - -C /

        else
            echo "No cache $cache_file_name found at $cache_bucket_url, skipping restore"
        fi
        ;;
    save)
        if [[ "$cache_file_exists" == "1" ]]; then
            echo "No cache $cache_file_name found at $cache_bucket_url"

            echo "Archiving cache ..."
            time tar cf - "${folders}" | pigz > "${cache_file_name}"

            echo "Uploading cache  ..."
            time gsutil cp "${cache_file_name}" "${cache_file_url}"
        else
            echo "Cache $cache_file_name at $cache_bucket_url already exists, skipping save"
        fi
        ;;
    *)
        echo $"Usage: $0 {restore|save}"
        exit 1
esac