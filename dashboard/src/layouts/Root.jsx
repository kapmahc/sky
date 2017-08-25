import React, {Component} from 'react'
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import {Layout, message} from 'antd'

import {TOKEN} from '../constants'
import {signIn, refresh} from '../actions'
import {get} from '../ajax'

class Widget extends Component {
  componentDidMount () {
    const {info, user, signIn, refresh} = this.props
    if (!user.uid) {
      const token = sessionStorage.getItem(TOKEN)
      if(token){
        signIn(token)
      }
    }
    if (info.languages.length === 0) {
      get('/site/info')
        .then(rst => {
          refresh(rst)
          document.title = `${rst.subTitle} | ${rst.title}`
        })
        .catch(message.error)
    }
  }
  render() {
    const {children} = this.props
    return (
      <Layout>
        {children}
      </Layout>
    );
  }
}

Widget.propTypes = {
  children: PropTypes.node.isRequired,
  signIn: PropTypes.func.isRequired,
  refresh: PropTypes.func.isRequired,
  user: PropTypes.object.isRequired,
  info: PropTypes.object.isRequired,
}


export default connect(
  state => ({
    user: state.currentUser,
    info: state.siteInfo,
  }),
  {signIn, refresh},
)(Widget)
