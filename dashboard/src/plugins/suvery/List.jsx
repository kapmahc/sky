import React, { Component } from 'react'
import { Row, Col, Button, message, Card } from 'antd'
import {FormattedMessage} from 'react-intl'
import { push } from 'react-router-redux'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'
import Markdown from 'react-markdown'


import Layout from '../../layouts/Application'
import {get} from '../../ajax'

class Widget extends Component {
  state = { items: []}
  componentDidMount () {
    get('/suvery').then(
      function (rst){
        this.setState({items: rst})
      }.bind(this)
    ).catch(message.error)
  }
  render() {
    const {push} = this.props

    return (
      <Layout breads={[{href: '/suvery', label: 'suvery.index.title'}]}>
        <Row gutter={16}>
          {this.state.items.map((f, i) => (<Col md={{span:8}} key={i}>
            <Card title={f.title}>
              <Markdown source={f.body}/>
              <p style={{paddingTop: "24px"}}>
                <Button type="primary" onClick={()=>push(`/suvery/apply/${f.id}`)}><FormattedMessage id="buttons.apply"/></Button>
                &nbsp;
                <Button type="danger" onClick={()=>push(`/suvery/cancel/${f.id}`)}><FormattedMessage id="buttons.cancel"/></Button>
              </p>
            </Card>
            <br/>
          </Col>))}
        </Row>
      </Layout>
    );
  }
}



Widget.propTypes = {
  push: PropTypes.func.isRequired,
}

export default connect(
  state => ({}),
  {push},
)(Widget)
