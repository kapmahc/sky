import React from 'react'
import { Route } from 'react-router'
import {FormattedMessage} from 'react-intl'

import Fail from '../../layouts/Fail'
import Home from './Home'
import Install from './Install'

const NotMatch = () => <Fail
  message={<FormattedMessage id="errors.not-match"/>}
  breadcrumbs={[]} />

export default [
  <Route key="site.home" exact path="/" component={Home}/>,
  <Route key="site.install" path="/install" component={Install}/>,

  <Route key="site.not-match" component={NotMatch}/>,
]
