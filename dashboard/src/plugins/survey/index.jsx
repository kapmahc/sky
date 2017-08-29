import React from 'react'
import { Route } from 'react-router'

import FormsEdit from './forms/Edit'
import FormsIndex from './forms/Index'

export default [
  <Route key="survey.new" path="/survey/forms/new" component={FormsEdit}/>,
  <Route key="survey.edit" path="/survey/forms/edit/:id" component={FormsEdit}/>,
  <Route key="survey.index" path="/survey/forms" component={FormsIndex}/>,
]
