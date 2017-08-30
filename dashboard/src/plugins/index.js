import auth from './auth'
import site from './site'
import survey from './survey'
import forum from './forum'

const routes = []
  .concat(auth)
  .concat(survey)
  .concat(forum)
  .concat(site)

export default routes
