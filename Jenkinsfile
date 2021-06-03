pipeline {
    agent any

    environment {
        REPOSITORY = "random_password_generator"
        TAG = "${env.BUILD_NUMBER}"
    }
    stages {        
        stage('docker build') {
            steps {
                withCredentials([usernamePassword(credentialsId: '528bdd1d-0b55-4ebd-996c-c5f27af385eb', passwordVariable: 'password', usernameVariable: 'username')]) {
                    echo 'Performing docker images build.'
                    sh 'docker build -t $REPOSITORY .'
                    echo 'Login and push images to Harbor'
                    sh 'docker login --username=$username --password=$password harbor.example.com'
                    sh 'docker tag $REPOSITORY harbor.example.com/test1/$REPOSITORY:$TAG'
                    sh 'docker push harbor.example.com/test1/$REPOSITORY:$TAG'
                }
            }
        }
        stage('clean images') {
            steps {
                echo 'clean images'
                sh 'docker rmi $REPOSITORY'
                sh 'docker rmi harbor.example.com/test1/random_password_generator:$TAG'
            }
        }
    }
}
