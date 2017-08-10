import React, { Component } from 'react'
import PropTypes from 'prop-types'
import {Card, Icon, Row, Col} from 'antd'
import {FormattedMessage} from 'react-intl'
import {connect} from 'react-redux'
import {Link} from 'react-router-dom'

import Application from './Application'
import {NonSignInLinks} from '../constants'

class Widget extends Component {
  render() {
    const {children, title} = this.props
    return (<Application>
        <Row>
          <Col md={{offset:8, span:8}}>
            <Card
              title={<FormattedMessage id={title}/>}
              extra={<a href="/" target="_blank"><FormattedMessage id="buttons.more"/></a>}>
              {children}
              <ul style={{marginTop: '20px'}}>
                {NonSignInLinks.map((l, i) => <li key={i}><Icon type={l.icon}/> <Link to={l.href}><FormattedMessage id={l.label}/></Link></li>)}
              </ul>
            </Card>
          </Col>
        </Row>
    </Application>)
  }
}

Widget.propTypes = {
  children: PropTypes.node.isRequired,
  title: PropTypes.string.isRequired,
}

export default connect(
  state => ({
    info: state.siteInfo,
  }),
  {},
)(Widget)
