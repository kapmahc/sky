import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import { Card, Col } from 'antd'
import {FormattedMessage} from 'react-intl'
import {Link} from 'react-router-dom'

import SignIn from '../auth/users/SignIn'
import Layout from '../../layouts/Dashboard'
import {dashboard} from '../../constants'

class Widget extends Component {
  render() {
    const {user} = this.props
    if(!user.uid){
      return <SignIn />
    }
    return <Layout breadcrumbs={[]}>
      {dashboard(user).map((m, i)=> <Col md={8} key={i}><Card title={<FormattedMessage id={m.label}/>}>
        {m.items.map((t, j)=> <p key={j}>
          <Link to={t.to}><FormattedMessage id={t.label}/></Link>
        </p>)}
      </Card><br/></Col>)}
    </Layout>
  }
}


Widget.propTypes = {
  user: PropTypes.object.isRequired,
}


export default connect(
  state => ({
    user: state.currentUser,
  }),
  {},
)(Widget)
