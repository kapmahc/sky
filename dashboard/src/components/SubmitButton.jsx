import React, { Component } from 'react'
import {FormattedMessage} from 'react-intl'
import {Button, Form} from 'antd'

const FormItem = Form.Item;

class Widget extends Component {
  render() {
    return (
      <FormItem>
        <Button type="primary" htmlType="submit" size="large">
          <FormattedMessage id="buttons.submit"/>
        </Button>
      </FormItem>
    );
  }
}

export default Widget;
