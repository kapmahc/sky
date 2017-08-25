import React, { Component } from 'react'
import { Form, Input, Col, message } from 'antd';
import {injectIntl, intlShape, FormattedMessage} from 'react-intl'

import Layout from '../../../layouts/Dashboard'
import SubmitButton from '../../../components/SubmitButton'
import {post, get} from '../../../ajax'

const FormItem = Form.Item;

class WidgetF extends Component {
  componentDidMount () {
    const {setFieldsValue} = this.props.form
    get('/users/info').then(
      function (rst){
        setFieldsValue({name: rst.name, email: rst.email})
      }
    ).catch(message.error)
  }
  handleSubmit = (e) => {
    e.preventDefault();
    const {formatMessage} = this.props.intl
    this.props.form.validateFieldsAndScroll((err, values) => {
     if (!err) {
       post('/users/info', values)
        .then((rst) => {
          message.success(formatMessage({id: 'messages.success'}))
        }).catch(message.error)
     }
    });
  }
  render() {
    const {formatMessage} = this.props.intl
    const { getFieldDecorator } = this.props.form;
    return (
      <Layout breadcrumbs={[{href: '/users/info', label:<FormattedMessage id='auth.users.info.title'/>}]}>
        <Col md={{span:8}}>
          <FormattedMessage id='auth.users.info.title' tagName='h2'/>
          <Form onSubmit={this.handleSubmit}>
            <FormItem
              label={<FormattedMessage id="attributes.email"/>}
            >
            {getFieldDecorator('email', {
              rules: [
              ],
            })(
              <Input disabled/>
            )}
            </FormItem>

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

            <SubmitButton />
          </Form>
        </Col>
      </Layout>
    );
  }
}


WidgetF.propTypes = {
  intl: intlShape.isRequired,
}

export default Form.create()(injectIntl(WidgetF))
