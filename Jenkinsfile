def components = ['products', 'reviews']

pipeline {
    agent none

    stages {
        stage ("build"){
            parallel {
                stage ("buid reviews") {
                    agent {
                        kubernetes {
                            name "reviews-builder"
                            yamlFile 'tools/did.yaml'
                        }
                    }
                    steps {
                        container("docker") {
                            sh """
                            echo reviews
                            docker version
                            """
                        }
                    }
                }
                stage ("buid products") {
                    agent {
                        kubernetes {
                            name "products-builder"
                            yamlFile 'tools/did.yaml'
                        }
                    }
                    steps {
                        container("docker") {
                            sh """
                            echo products
                            docker version
                            """
                        }
                    }
                }
            }
        }
    }
}