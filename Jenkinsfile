pipeline {
    agent { 
        docker {
            image 'golang:1.20'
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
                sh 'go run main.go'
                echo "doing test stuff.."
                '''
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