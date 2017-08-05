import registerServiceWorker from './registerServiceWorker'
import {detectLocale} from './intl'
import main from './main'

const user = detectLocale()
main('root', user)
registerServiceWorker()
