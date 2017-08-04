import auth from './auth'
import site from './site'
import suvery from './suvery'
// import mall from './mall'

const routes = []
  .concat(auth)
  // .concat(mall)
  .concat(suvery)
  .concat(site)

export default routes
