def components = ['products', 'reviews']

pipeline {
    agent none

    stages {
        stage ("build"){
            parallel {
                for (component in components) {
                    stage ("buid ${component}") {
                        agent {
                            kubernetes {
                                name "${component}-builder"
                                yamlFile 'tools/did.yaml'
                            }
                        }
                        steps {
                            container("docker") {
                                sh """
                                echo ${component}
                                docker version
                                """
                            }
                        }
                    }
                }
            }
        }
    }
}