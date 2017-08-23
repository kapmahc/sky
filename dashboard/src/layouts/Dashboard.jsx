import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {connect} from 'react-redux'
import {FormattedMessage} from 'react-intl'
import {Layout, Menu, Row} from 'antd'

import Root from './Root'
import Fail from './Fail'
import Footer from '../components/Footer'
import Breadcrumb from '../components/Breadcrumb'

const { Header, Content } = Layout

class Widget extends Component {
  render() {
    const {children, info, user, breadcrumbs} = this.props
    if (!user.uid) {
      return <Fail
        message={<FormattedMessage id="errors.not-allow"/>}
        breadcrumbs={[
          {label:<FormattedMessage id="auth.users.sign-in.title"/>, href: "/users/sign-in"},
        ]} />
    }
    return (<Root>
      <Header>
        <div className="logo" />
        <Menu
          theme="dark"
          mode="horizontal"
          defaultSelectedKeys={[]}
          style={{ lineHeight: '64px' }}
        >
          <Menu.Item key="home">{info.title}</Menu.Item>
        </Menu>
      </Header>
      <Content style={{ padding: '0 50px'}}>
        <Breadcrumb items={breadcrumbs}/>
        <Row style={{ background: '#fff', padding: 24, minHeight: 380 }}>
          {children}
        </Row>
      </Content>
      <Footer/>
    </Root>)
  }
}

Widget.propTypes = {
  children: PropTypes.node.isRequired,
  info: PropTypes.object.isRequired,
  user: PropTypes.object.isRequired,
  breadcrumbs: PropTypes.array.isRequired,
}

export default connect(
  state => ({
    info: state.siteInfo,
    user: state.currentUser,
  }),
  {},
)(Widget)
