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
import {TOKEN} from '../constants'
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
        break;
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

    var sider = [
      {key: 'personal'},
      {}
    ]
    if(user.admin){
      sider.push({
        key: ''
      })
    }
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
          <SubMenu key="personal" title={<span><Icon type="user" /><FormattedMessage id="sider.personal" values={user}/></span>}>
            <Menu.Item key="to-/users/logs"><FormattedMessage id="auth.users.logs.title"/></Menu.Item>
            <Menu.Item key="to-/users/change-password"><FormattedMessage id="auth.users.change-password.title"/></Menu.Item>
            <Menu.Item key="to-/users/info"><FormattedMessage id="auth.users.info.title"/></Menu.Item>
            <Menu.Item key="to-/attachments"><FormattedMessage id="auth.attachments.index.title"/></Menu.Item>
          </SubMenu>
          <SubMenu key="site" title={<span><Icon type="setting" /><FormattedMessage id="sider.site" values={user}/></span>}>
            <Menu.Item key="to-/admin/status"><FormattedMessage id="site.admin.status.title"/></Menu.Item>
            <Menu.Item key="to-/admin/site/info"><FormattedMessage id="site.admin.site.info.title"/></Menu.Item>
            <Menu.Item key="to-/admin/site/author"><FormattedMessage id="site.admin.site.author.title"/></Menu.Item>
            <Menu.Item key="to-/admin/seo"><FormattedMessage id="site.admin.seo.title"/></Menu.Item>
            <Menu.Item key="to-/admin/smtp"><FormattedMessage id="site.admin.smtp.title"/></Menu.Item>
            <Menu.Item key="to-/admin/paypal"><FormattedMessage id="site.admin.paypal.title"/></Menu.Item>
            <Menu.Item key="to-/admin/locales"><FormattedMessage id="site.admin.locales.index.title"/></Menu.Item>
            <Menu.Item key="to-/admin/links"><FormattedMessage id="site.admin.links.index.title"/></Menu.Item>
            <Menu.Item key="to-/admin/cards"><FormattedMessage id="site.admin.cards.index.title"/></Menu.Item>
            <Menu.Item key="to-/admin/users"><FormattedMessage id="site.admin.users.index.title"/></Menu.Item>
            <Menu.Item key="to-/admin/friend-links"><FormattedMessage id="site.admin.friend-links.index.title"/></Menu.Item>
            <Menu.Item key="to-/leave-words"><FormattedMessage id="site.leave-words.index.title"/></Menu.Item>
          </SubMenu>
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
            <Row>
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
