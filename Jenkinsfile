pipeline {
  agent any
  tools {
    go "Go 1.13"
  }
  options {
    checkoutToSubdirectory('src/github.com/pikachu')
  }
  environment {
    GOPATH = "$WORKSPACE"
    DIRECTORY = "src/github.com/pikachu"
  }
  stages {
    stage("Lint") {
      steps {
        dir("$DIRECTORY") {
          sh "make fmt && git diff --exit-code"
        }
      }
    }
    stage("Test") {
      steps {
        dir("$DIRECTORY") {
          sh "make test"
        }
      }
    }
    stage("Build") {
      steps {
        withDockerRegistry([credentialsId: "<insert-the-creds-id>", url: ""]) {
          dir("$DIRECTORY") {
            sh "make docker push"
          }
        }
      }
    }
    stage("Push") {
      when {
        branch "master"
      }
      steps {
        withDockerRegistry([credentialsId: "<insert-the-creds-id>", url: ""]) {
          dir("$DIRECTORY") {
            sh "make push IMAGE_VERSION=latest"
          }
        }
      }
    }
    
  }
  post {
    always {
      dir("$DIRECTORY") {
        sh "make clean || true"
      }
      cleanWs()
    }
  }
}
