pipeline {
    agent any
    tools {
        go 'go1.14'
    }
    stages {        
        stage('Pre Test') {
            steps {
                go version
            }
        }
        
        stage('Build') {
            steps {
                echo 'Compiling and building'
                go build
            }
        }

        stage('Test') {
            steps {
                ./hello-world
            }
        }
        
    }
    post {
        always {
            echo 'Finish Pipeline'
        }
    }  
}