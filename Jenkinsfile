pipeline {
    agent {
        docker{
            image 'golang:latest'
        }
    }
    
    // triggers {
    //     pollSCM '*/5 * * * *'
    // }

    stages {
        stage('Build') {
            steps {
                echo "Building.."
                sh '''
                echo "doing build stuff.."
                '''
            }
        }
        stage('Test') {
            steps {
                echo "Testing.."
                sh 'bash -c "go run main.go"'
                echo "doing test stuff.."
                
            }
        }
        stage('Deliver') {
            steps {
                echo 'Deliver....'
                sh '''
                echo "doing delivery stuff.."
                '''
            }
        }
    }
}