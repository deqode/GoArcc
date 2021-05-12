const {PHASE_DEVELOPMENT_SERVER,PHASE_PRODUCTION_SERVER}=require ("next/constants")
module.exports =(phase)=> {
  if(phase===PHASE_PRODUCTION_SERVER){
    return{
      env: {
        SERVER: 'http://localhost:8082/v1',
        CENTRIFUGO: 'ws://localhost:7070/connection/websocket',
        GQLHTTP: 'http://localhost:8081/graphql',
        SESSION_PASSWORD:'complex_password_at_least_32_characters_long',
        COOKIE_NAME:'alfred'
      },
    }
  }
  if(phase===PHASE_DEVELOPMENT_SERVER){
    return{
      env: {
        SERVER: 'http://localhost:8082/v1',
        CENTRIFUGO: 'ws://localhost:7070/connection/websocket',
        GQLHTTP: 'http://localhost:8081/graphql',
        SESSION_PASSWORD:'complex_password_at_least_32_characters_long',
        COOKIE_NAME:'alfred'
      },
    }
  }
  //default values
  return {}
}
