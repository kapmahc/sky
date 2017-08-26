import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {connect} from 'react-redux'

import SignIn from '../auth/users/SignIn'
import Layout from '../../layouts/Dashboard'

class Widget extends Component {
  render() {
    const {user} = this.props
    if(!user.uid){
      return <SignIn />
    }
    return <Layout breadcrumbs={[]}>
      <h1>Help</h1>
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
