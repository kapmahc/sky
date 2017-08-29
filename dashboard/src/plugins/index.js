import auth from './auth'
import site from './site'
import survey from './survey'

const routes = []
  .concat(auth)
  .concat(survey)
  .concat(site)

export default routes
