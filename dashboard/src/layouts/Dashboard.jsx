import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import {FormattedMessage, injectIntl, intlShape} from 'react-intl'
import {Layout, Menu, Row, Icon, message, Modal} from 'antd'
import { push } from 'react-router-redux'

import Root from './Root'
import Fail from './Fail'
import Footer from '../components/Footer'
import Breadcrumb from '../components/Breadcrumb'
import {TOKEN, dashboard} from '../constants'
import {signOut} from '../actions'
import {_delete} from '../ajax'

const { Sider, Content } = Layout
const SubMenu = Menu.SubMenu
const confirm = Modal.confirm

class Widget extends Component {
  onSiderMenu = ({key}) => {
    const {push, signOut} = this.props
    const {formatMessage} = this.props.intl
    const to = 'to-'
    if(key.startsWith(to)){
      push(key.substring(to.length))
      return
    }

    switch (key) {
      case 'sign-out':
          confirm({
            title: formatMessage({id: "messages.are-you-sure"}),
            onOk() {
              _delete('/users/sign-out').then((rst) => {
                signOut()
                sessionStorage.removeItem(TOKEN)
                push('/users/sign-in')
                message.success(formatMessage({id: 'messages.success'}))
              }).catch(message.error)
            }
          });
        return
      case 'home':
        window.open('/', '_blank')
        return
      default:
        console.log(key)
    }
  }
  render() {
    const {children, admin, user, breadcrumbs} = this.props
    // console.log(user)
    if (!user.uid || (admin && !user.admin)) {
      return <Fail
        message={<FormattedMessage id="errors.not-allow"/>}
        breadcrumbs={[
          {label:<FormattedMessage id="auth.users.sign-in.title"/>, href: "/users/sign-in"},
        ]} />
    }

    var sider = dashboard(user)
    return (<Root>
      <Sider
        breakpoint="lg"
        collapsedWidth="0"
        onCollapse={(collapsed, type) => { console.log(collapsed, type); }}
      >
        <div className="logo" />
        <Menu theme="dark" mode="inline" onClick={this.onSiderMenu} defaultSelectedKeys={[]}>
          <Menu.Item key="home">
            <Icon type="home" />
            <FormattedMessage id="sider.home" className="nav-text"/>
          </Menu.Item>
          {sider.map((m, i)=><SubMenu key={`sider-${i}`} title={<span><Icon type={m.icon} /><FormattedMessage id={m.label}/></span>}>
            {m.items.map((t)=><Menu.Item key={`to-${t.to}`}><FormattedMessage id={t.label}/></Menu.Item>)}
          </SubMenu>)}
          <Menu.Item key="sign-out">
            <Icon type="logout" />
            <FormattedMessage id="sider.sign-out" className="nav-text"/>
          </Menu.Item>
        </Menu>
      </Sider>
      <Layout>
        <Content style={{ padding: '0 50px'}}>
          <Breadcrumb items={breadcrumbs}/>
          <div style={{ padding: 24, background: '#fff', minHeight: 380 }}>
            <Row gutter={16}>
              {children}
            </Row>
          </div>
        </Content>
        <Footer />
      </Layout>
    </Root>)
  }
}

Widget.propTypes = {
  children: PropTypes.node.isRequired,
  user: PropTypes.object.isRequired,
  push: PropTypes.func.isRequired,
  signOut: PropTypes.func.isRequired,
  breadcrumbs: PropTypes.array.isRequired,
  intl: intlShape.isRequired,
}

const WidgetM = injectIntl(Widget)

export default connect(
  state => ({
    info: state.siteInfo,
    user: state.currentUser,
  }),
  {push, signOut},
)(WidgetM)
