import React, { Component } from 'react'
import { Form, Input, message } from 'antd';
import {injectIntl, intlShape, FormattedMessage} from 'react-intl'
import { push } from 'react-router-redux'
import { connect } from 'react-redux'
import PropTypes from 'prop-types'

import Layout from '../../../layouts/Dashboard'
import SubmitButton from '../../../components/SubmitButton'
import {post, get} from '../../../ajax'

const FormItem = Form.Item

class WidgetF extends Component {
  componentDidMount () {
    const {setFieldsValue} = this.props.form
    const {id} = this.props.match.params
    if (id) {
      get(`/forum/tags/${id}`).then((rst)=>setFieldsValue({name: rst.name})).catch(message.error)
    }
  }
  handleSubmit = (e) => {
    e.preventDefault();
    const {formatMessage} = this.props.intl
    const {push, match} = this.props
    const {id} = match.params

    this.props.form.validateFieldsAndScroll((err, values) => {
     if (!err) {
       post(id ? `/forum/tags/${id}` : '/forum/tags', Object.assign({}, values))
        .then((rst) => {
          message.success(formatMessage({id: 'messages.success'}))
          push('/forum/tags')
        }).catch(message.error)
     }
    });
  }
  render() {
    const {formatMessage} = this.props.intl
    const { getFieldDecorator } = this.props.form;
    const {id} = this.props.match.params
    return (
      <Layout admin breadcrumbs={[
          {href: '/forum/tags', label: <FormattedMessage id='forum.tags.index.title'/>},
          {href: id ? `/forum/tags/edit/${id}` : '/forums/tags/new', label: <FormattedMessage id={id ? 'buttons.edit': 'buttons.new'}/>},
        ]}>
        <Form onSubmit={this.handleSubmit}>
          <FormItem
            label={<FormattedMessage id="attributes.name"/>}
            hasFeedback
          >
          {getFieldDecorator('name', {
            rules: [{ required: true, message: formatMessage({id:"errors.not-empty"})}],
          })(
            <Input />
          )}
          </FormItem>

          <SubmitButton />
        </Form>
      </Layout>
    );
  }
}


WidgetF.propTypes = {
  intl: intlShape.isRequired,
  push: PropTypes.func.isRequired,
}

const Widget = Form.create()(injectIntl(WidgetF))

export default connect(
  state => ({}),
  {push},
)(Widget)
