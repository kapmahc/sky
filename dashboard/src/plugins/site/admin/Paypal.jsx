import React, { Component } from 'react'
import { Form, Input, message } from 'antd';
import {injectIntl, intlShape, FormattedMessage} from 'react-intl'

import Layout from '../../../layouts/Dashboard'
import SubmitButton from '../../../components/SubmitButton'
import {post, get} from '../../../ajax'

const FormItem = Form.Item;

class WidgetF extends Component {
  componentDidMount () {
    const {setFieldsValue} = this.props.form
    get('/admin/paypal').then(
      function (rst){
        setFieldsValue({
          donate: rst.donate,
        })
      }
    ).catch(message.error)
  }
  handleSubmit = (e) => {
    e.preventDefault();
    const {formatMessage} = this.props.intl
    this.props.form.validateFieldsAndScroll((err, values) => {
     if (!err) {
       post('/admin/paypal', values)
        .then((rst) => {
          message.success(formatMessage({id: 'messages.success'}))
        }).catch(message.error)
     }
    });
  }
  render() {
    const { getFieldDecorator } = this.props.form;
    return (
      <Layout admin breadcrumbs={[{href: '/admin/paypal', label: <FormattedMessage id='site.admin.paypal.title'/>}]}>
        <Form onSubmit={this.handleSubmit}>

          <FormItem

            label={<FormattedMessage id="site.admin.paypal.donate-form"/>}
            hasFeedback
          >
          {getFieldDecorator('donate', {
            rules: [],
          })(
            <Input type="textarea" rows={10} />
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
}

export default Form.create()(injectIntl(WidgetF))
