import React from 'react'
import { Route } from 'react-router'

import Manage from './Manage'
import Edit from './Edit'
import Show from './Show'
import Index from './List'

const Apply = ({match}) => <Show action="apply" id={match.params.id} />
const Cancel = ({match}) => <Show action="cancel" id={match.params.id}/>

export default [
  <Route key="suvery.manage" path="/suvery/manage" component={Manage}/>,
  <Route key="suvery.new" path="/suvery/new" component={Edit}/>,
  <Route key="suvery.edit" path="/suvery/edit/:id" component={Edit}/>,
  <Route key="suvery.apply" path="/suvery/apply/:id" component={Apply}/>,
  <Route key="suvery.cancel" path="/suvery/cancel/:id" component={Cancel}/>,
  <Route key="suvery.index" path="/suvery" component={Index}/>,
]
