import React from 'react'
import { Route } from 'react-router'
import {FormattedMessage} from 'react-intl'

import Fail from '../../layouts/Fail'
import Home from './Home'
import Install from './Install'

import LeaveWordsNew from './leave-words/New'
import LeaveWordsIndex from './leave-words/Index'

import AdminSiteInfo from './admin/site/Info'
import AdminSiteAuthor from './admin/site/Author'

import AdminSeo from './admin/Seo'
import AdminSmtp from './admin/Smtp'
import AdminStatus from './admin/Status'
import AdminPaypal from './admin/Paypal'

import AdminUsersIndex from './admin/users/Index'
import AdminLocalesIndex from './admin/locales/Index'
import AdminLocalesEdit from './admin/locales/Edit'

import AdminLinksIndex from './admin/links/Index'
import AdminLinksEdit from './admin/links/Edit'

import AdminCardsIndex from './admin/cards/Index'
import AdminCardsEdit from './admin/cards/Edit'

import AdminFriendLinksIndex from './admin/friend-links/Index'
import AdminFriendLinksEdit from './admin/friend-links/Edit'

const NotMatch = () => <Fail
  message={<FormattedMessage id="errors.not-match"/>}
  breadcrumbs={[]} />

export default [
  <Route key="site.home" exact path="/" component={Home}/>,
  <Route key="site.install" path="/install" component={Install}/>,

  <Route key="site.leave-words.index" path="/leave-words" component={LeaveWordsIndex}/>,
  <Route key="site.leave-words.new" path="/leave-words/new" component={LeaveWordsNew}/>,

  <Route key="site.admin.site.info" path="/admin/site/info" component={AdminSiteInfo}/>,
  <Route key="site.admin.site.author" path="/admin/site/author" component={AdminSiteAuthor}/>,

  <Route key="site.admin.smtp" path="/admin/smtp" component={AdminSmtp}/>,
  <Route key="site.admin.seo" path="/admin/seo" component={AdminSeo}/>,
  <Route key="site.admin.status" path="/admin/status" component={AdminStatus}/>,
  <Route key="site.admin.paypal" path="/admin/paypal" component={AdminPaypal}/>,

  <Route key="site.admin.users.index" path="/admin/users" component={AdminUsersIndex}/>,
  <Route key="site.admin.locales.new" path="/admin/locales/new" component={AdminLocalesEdit}/>,
  <Route key="site.admin.locales.edit" path="/admin/locales/edit/:code" component={AdminLocalesEdit}/>,
  <Route key="site.admin.locales.index" path="/admin/locales" component={AdminLocalesIndex}/>,

  <Route key="site.admin.links.new" path="/admin/links/new" component={AdminLinksEdit}/>,
  <Route key="site.admin.links.edit" path="/admin/links/edit/:id" component={AdminLinksEdit}/>,
  <Route key="site.admin.links.index" path="/admin/links" component={AdminLinksIndex}/>,

  <Route key="site.admin.friend-links.new" path="/admin/friend-links/new" component={AdminFriendLinksEdit}/>,
  <Route key="site.admin.friend-links.edit" path="/admin/friend-links/edit/:id" component={AdminFriendLinksEdit}/>,
  <Route key="site.admin.friend-links.index" path="/admin/friend-links" component={AdminFriendLinksIndex}/>,

  <Route key="site.admin.cards.new" path="/admin/cards/new" component={AdminCardsEdit}/>,
  <Route key="site.admin.cards.edit" path="/admin/cards/edit/:id" component={AdminCardsEdit}/>,
  <Route key="site.admin.cards.index" path="/admin/cards" component={AdminCardsIndex}/>,

  <Route key="site.not-match" component={NotMatch}/>,
]
