import React, { Component } from 'react'
import PropTypes from 'prop-types'
import { Layout, Menu, Breadcrumb, Icon, Modal, message } from 'antd';
import {injectIntl, intlShape, FormattedMessage} from 'react-intl'
import { connect } from 'react-redux'
import { push } from 'react-router-redux'
import {Link} from 'react-router-dom'

import {NonSignInLinks, TOKEN} from '../constants'
import {setLocale} from '../intl'
import {signIn, signOut, refresh} from '../actions'
import {get, _delete} from '../ajax'
import menus from '../menus'

const { SubMenu } = Menu;
const { Content, Sider} = Layout;
const confirm = Modal.confirm;

class WidgetF extends Component {
  handleMenu = ({key}) => {
    const {push, signOut} = this.props
    const {formatMessage} = this.props.intl

    const to = 'to-'
    if(key.startsWith(to)){
      push(key.substring(to.length))
      return
    }

    const lng = 'language-'
    if(key.startsWith(lng)){
      setLocale(key.substring(lng.length))
      return
    }

    switch (key) {
      case 'auth.users.sign-out':
        confirm({
          title: formatMessage({id: "messages.are-you-sure"}),
          onOk() {
            _delete('/users/sign-out').then((rst) => {
              signOut()
              sessionStorage.removeItem(TOKEN)
              push('/')
              message.success(formatMessage({id: 'messages.success'}))
            }).catch(message.error)
          }
        });
        break;
      default:
        console.error(key)
    }
  }
  componentDidMount () {
    const {info, user, signIn, refresh} = this.props
    if (!user.uid) {
      const token = sessionStorage.getItem(TOKEN)
      if(token){
        signIn(token)
      }
    }
    if (info.languages.length === 0) {
      get('/api/site/info')
        .then(rst => {
          refresh(rst)
          document.title = `${rst.subTitle} | ${rst.title}`
        })
        .catch(message.error)
    }
  }
  render() {
    const {children, user, info, breads} = this.props
    return (
      <div>
        <h1>header</h1>
        {children}
      </div>
    );
  }
}

WidgetF.propTypes = {
  children: PropTypes.node.isRequired,
  push: PropTypes.func.isRequired,
  refresh: PropTypes.func.isRequired,
  signIn: PropTypes.func.isRequired,
  signOut: PropTypes.func.isRequired,
  user: PropTypes.object.isRequired,
  info: PropTypes.object.isRequired,
  breads: PropTypes.array,
  mustSignIn: PropTypes.bool,
  mustAdmin: PropTypes.bool,
  intl: intlShape.isRequired,
}


const Widget = injectIntl(WidgetF)

export default connect(
  state => ({
    user: state.currentUser,
    info: state.siteInfo,
  }),
  {push, signIn, refresh, signOut},
)(Widget)
