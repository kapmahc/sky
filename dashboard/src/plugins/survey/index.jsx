import React from 'react'
import { Route } from 'react-router'

import FormsEdit from './forms/Edit'
import FormsIndex from './forms/Index'
// import ReportsIndex from './reports/Index'

export default [
  <Route key="survey.forms.new" path="/survey/forms/new" component={FormsEdit}/>,
  <Route key="survey.forms.edit" path="/survey/forms/edit/:id" component={FormsEdit}/>,
  <Route key="survey.forms.index" path="/survey/forms" component={FormsIndex}/>,

  // <Route key="survey.reports.index" path="/survey/reports" component={ReportsIndex}/>,
]
