import React from 'react'
import { Route } from 'react-router'

import ArticlesEdit from './articles/Edit'
import ArticlesIndex from './articles/Index'
import TagsEdit from './tags/Edit'
import TagsIndex from './tags/Index'
import CommentsEdit from './comments/Edit'
import CommentsIndex from './comments/Index'

export default [
  <Route key="forum.articles.new" path="/forum/articles/new" component={ArticlesEdit}/>,
  <Route key="forum.articles.edit" path="/forum/articles/edit/:id" component={ArticlesEdit}/>,
  <Route key="forum.articles.index" path="/forum/articles" component={ArticlesIndex}/>,

  <Route key="forum.tags.new" path="/forum/tags/new" component={TagsEdit}/>,
  <Route key="forum.tags.edit" path="/forum/tags/edit/:id" component={TagsEdit}/>,
  <Route key="forum.tags.index" path="/forum/tags" component={TagsIndex}/>,

  <Route key="forum.comments.new" path="/forum/comments/new" component={CommentsEdit}/>,
  <Route key="forum.comments.edit" path="/forum/comments/edit/:id" component={CommentsEdit}/>,
  <Route key="forum.comments.index" path="/forum/comments" component={CommentsIndex}/>,

]
