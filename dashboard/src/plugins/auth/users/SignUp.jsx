import React, { Component } from 'react';
import PropTypes from 'prop-types'
import { Form, Input, message } from 'antd';
import {injectIntl, intlShape, FormattedMessage} from 'react-intl'
import { connect } from 'react-redux'
import { push } from 'react-router-redux'

import Layout from './Layout'
import SubmitButton from '../../../components/SubmitButton'
import {post} from '../../../ajax'

const FormItem = Form.Item;

class WidgetF extends Component {
  handleSubmit = (e) => {
    e.preventDefault();
    const {formatMessage} = this.props.intl
    const {push} = this.props
    this.props.form.validateFieldsAndScroll((err, values) => {
     if (!err) {
       post('/users/sign-up', values)
        .then((rst) => {
          message.success(formatMessage({id: 'auth.users.confirm.success'}))
          push('/users/sign-in')
        }).catch(message.error)
     }
    });
  }
  checkPasswords = (rule, value, callback) => {
    const {formatMessage} = this.props.intl
    const form = this.props.form;
    if (value && value !== form.getFieldValue('password')) {
      callback(formatMessage({id: "errors.passwordConfirmation"}));
    } else {
      callback();
    }
  }
  render() {
    const {formatMessage} = this.props.intl
    const { getFieldDecorator } = this.props.form;
    return (
      <Layout href="/users/sign-up" title="auth.users.sign-up.title">
        <Form onSubmit={this.handleSubmit}>
          <FormItem
            label={<FormattedMessage id="attributes.username"/>}
            hasFeedback
          >
          {getFieldDecorator('name', {
            rules: [{ required: true, message: formatMessage({id:"errors.not-empty"})}],
          })(
            <Input />
          )}
          </FormItem>

          <FormItem
            label={<FormattedMessage id="attributes.email"/>}
            hasFeedback
          >
          {getFieldDecorator('email', {
            rules: [
              { required: true, message: formatMessage({id:"errors.not-empty"})},
              { type: 'email', message: formatMessage({id:"errors.not-valid-email"})},
            ],
          })(
            <Input />
          )}
          </FormItem>

          <FormItem
            label={<FormattedMessage id="attributes.password"/>}
            hasFeedback
          >
          {getFieldDecorator('password', {
            rules: [
              { required: true, min: 6, max: 32, message: formatMessage({id:"errors.password"})},
            ],
          })(
            <Input type="password" />
          )}
          </FormItem>

          <FormItem
            label={<FormattedMessage id="attributes.passwordConfirmation"/>}
            hasFeedback
          >
          {getFieldDecorator('passwordConfirmation', {
            rules: [
              { required: true, message: formatMessage({id:"errors.not-empty"})},
              {validator: this.checkPasswords},
            ],
          })(
            <Input type="password" />
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
