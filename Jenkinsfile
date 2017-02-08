properties([
  buildDiscarder(
    logRotator(
      artifactDaysToKeepStr: '14',
      artifactNumToKeepStr: '20',
      daysToKeepStr: '28',
      numToKeepStr: '20'
    )
  ),
])
node ('linux'){
  timestamps {
    withEnv(["GIT_BRANCH=${env.BRANCH_NAME}"]) {
      stage 'Prepare'
      checkout scm
      stage 'Build'
      sh "make -e package"
      stage 'Test'
      sh "make -e test"
      stage 'Publish'
      sh "make -e push"
      slackSend channel: '#slacktesting', color: 'good', message: 'built webhealth', teamDomain: 'dronedeploy'
      if ("stage".equals(env.BRANCH_NAME)) {
        sh "make -e smoke-all"
      }
    }
  }
}
