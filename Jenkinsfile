def components = ['products', 'reviews']

pipeline {
    agent none
    stages {
        stage ("build"){
            parallel {
                stage ("buid reviews") {
                    agent {
                        kubernetes {
                            label "reviews-builder"
                            yamlFile 'tools/did.yaml'
                        }
                    }
                    steps {
                        container("docker") {
                            sh """
                            docker build \
                                -f ./services/reviews/Dockerfile \
                                --build-arg PROJECT=./services/reviews \
                                -t reviews:${env.GIT_COMMIT} .
                            """
                        }
                    }
                }
                stage ("buid products") {
                    agent {
                        kubernetes {
                            label "products-builder"
                            yamlFile 'tools/did.yaml'
                        }
                    }
                    steps {
                        container("docker") {
                            sh """
                            docker build \
                                -f ./services/products/Dockerfile \
                                --build-arg PROJECT=./services/reviews \
                                -t products:${env.GIT_COMMIT} .
                            """
                        }
                    }
                }
            }
        }
    }
}