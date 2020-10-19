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
                go run hello-world.go
            }
        }
        
    }
    post {
        always {
            echo 'Finish Pipeline'
        }
    }  
}