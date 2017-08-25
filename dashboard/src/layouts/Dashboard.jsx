import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import {FormattedMessage} from 'react-intl'
import {Layout, Menu, Row, Icon} from 'antd'
import { push } from 'react-router-redux'

import Root from './Root'
import Fail from './Fail'
import Footer from '../components/Footer'
import Breadcrumb from '../components/Breadcrumb'

const { Sider, Content } = Layout
const SubMenu = Menu.SubMenu

class Widget extends Component {
  onSiderMenu = ({key}) => {
    const {push} = this.props
    const to = 'to-'
    if(key.startsWith(to)){
      push(key.substring(to.length))
      return
    }
    console.log(key)
  }
  render() {
    const {children, user, breadcrumbs} = this.props
    if (!user.uid) {
      return <Fail
        message={<FormattedMessage id="errors.not-allow"/>}
        breadcrumbs={[
          {label:<FormattedMessage id="auth.users.sign-in.title"/>, href: "/users/sign-in"},
        ]} />
    }
    return (<Root>
      <Sider
        breakpoint="lg"
        collapsedWidth="0"
        onCollapse={(collapsed, type) => { console.log(collapsed, type); }}
      >
        <div className="logo" />
        <Menu theme="dark" mode="inline" onClick={this.onSiderMenu} defaultSelectedKeys={[]}>
          <Menu.Item key="site.home">
            <Icon type="home" />
            <FormattedMessage id="sider.home.title" className="nav-text"/>
          </Menu.Item>
          <SubMenu key="personal" title={<span><Icon type="user" /><FormattedMessage id="sider.personal.title" values={user}/></span>}>
            <Menu.Item key="to-/users/logs"><FormattedMessage id="auth.users.logs.title"/></Menu.Item>
            <Menu.Item key="to-/users/change-password"><FormattedMessage id="auth.users.change-password.title"/></Menu.Item>
            <Menu.Item key="to-/users/info"><FormattedMessage id="auth.users.info.title"/></Menu.Item>
            <Menu.Item key="to-/attachments"><FormattedMessage id="auth.attachments.index.title"/></Menu.Item>
          </SubMenu>
          <Menu.Item key="1">
            <Icon type="user" />
            <span className="nav-text">nav 1</span>
          </Menu.Item>
          <Menu.Item key="2">
            <Icon type="video-camera" />
            <span className="nav-text">nav 2</span>
          </Menu.Item>
          <Menu.Item key="3">
            <Icon type="upload" />
            <span className="nav-text">nav 3</span>
          </Menu.Item>
          <Menu.Item key="4">
            <Icon type="user" />
            <span className="nav-text">nav 4</span>
          </Menu.Item>
          <Menu.Item key="5">
            <Icon type="user" />
            <span className="nav-text">nav 4</span>
          </Menu.Item>
         <SubMenu key="sub2" title={<span><Icon type="appstore" /><span>Navigation Two</span></span>}>
           <Menu.Item key="9">Option 9</Menu.Item>
           <Menu.Item key="10">Option 10</Menu.Item>
           <SubMenu key="sub3" title="Submenu">
             <Menu.Item key="11">Option 11</Menu.Item>
             <Menu.Item key="12">Option 12</Menu.Item>
           </SubMenu>
         </SubMenu>
          <Menu.Item key="6">
            <Icon type="user" />
            <span className="nav-text">nav 4</span>
          </Menu.Item>
          <Menu.Item key="7">
            <Icon type="user" />
            <span className="nav-text">nav 4</span>
          </Menu.Item>
          <Menu.Item key="8">
            <Icon type="user" />
            <span className="nav-text">nav 4</span>
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
  breadcrumbs: PropTypes.array.isRequired,
}

export default connect(
  state => ({
    info: state.siteInfo,
    user: state.currentUser,
  }),
  {push},
)(Widget)
