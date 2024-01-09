pipeline {
    agent any

    tools {
        go 'go 1.21.5'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout([
                    $class: 'GitSCM', 
                    branches: [[name: '*/main']], 
                    doGenerateSubmoduleConfigurations: false, 
                    extensions: [], 
                    submoduleCfg: [], 
                    userRemoteConfigs: [[url: 'https://github.com/tkhv/GoRedis.git']]
                ])
            }
        }

        stage('Clear Workspace') {
            steps {
                deleteDir()
            }
        }
        
        stage('Build') {
            steps {
                sh 'git clone https://github.com/tkhv/GoRedis.git .'
                sh 'go build'
            }
        }

        stage('Deploy') {
            steps {
                sshPublisher(
                    continueOnError: false, 
                    publishers: [
                        sshPublisherDesc(
                            configName: 'dockerhost', 
                            transfers: [
                                sshTransfer(
                                    execCommand: 'docker stop $(docker ps -a -q)', 
                                    execTimeout: 120000
                                ),
                                sshTransfer(
                                    sourceFiles: '*', 
                                    removePrefix: '', 
                                    remoteDirectory: '', 
                                    execCommand: 'docker build -t goredis .', 
                                    execTimeout: 120000
                                ),
                                sshTransfer(
                                    execCommand: 'docker run -d -p 6379:6379 goredis', 
                                    execTimeout: 120000
                                )
                            ]
                        )
                    ]
                )
            }
        }
    }
}