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
                            docker build -t reviews:${env.GIT_COMMIT} .
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
                            docker build -t products:${env.GIT_COMMIT} .
                            """
                        }
                    }
                }
            }
        }
    }
}