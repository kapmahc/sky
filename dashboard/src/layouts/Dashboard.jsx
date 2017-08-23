import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {Col, Icon} from 'antd'
import {FormattedMessage} from 'react-intl'
import {Link} from 'react-router-dom'

import Layout from './Application'

class Widget extends Component {
  render() {
    const {children, breadcrumbs} = this.props
    const links = [
      {icon: 'login', href: '/users/sign-in', label: 'auth.users.sign-in.title'},
      {icon: 'user-add', href: '/users/sign-up', label: 'auth.users.sign-up.title'},
      {icon: 'retweet', href: '/users/forgot-password', label: 'auth.users.forgot-password.title'},
      {icon: 'check-circle-o', href: '/users/confirm', label: 'auth.users.confirm.title'},
      {icon: 'unlock', href: '/users/unlock', label: 'auth.users.unlock.title'},
      {icon: 'question-circle-o', href: '/leave-words/new', label: 'site.leave-words.new.title'},
    ]
    return <Layout breadcrumbs={[
        {label: <FormattedMessage id={title}/>, href}
      ]}>
      <Col md={{offset:8, span:8}}>
        <FormattedMessage tagName="h2" id={title} />
        <div style={{marginTop: '20px'}}>
          {children}
        </div>
        <ul style={{marginTop: '20px'}}>
          {links.map((l, i) => <li key={i}><Icon type={l.icon}/> <Link to={l.href}><FormattedMessage id={l.label}/></Link></li>)}
        </ul>
      </Col>
    </Layout>
  }
}


Widget.propTypes = {
  user: PropTypes.object.isRequired,
  breadcrumbs: PropTypes.array.isRequired,
}

export default Widget;
