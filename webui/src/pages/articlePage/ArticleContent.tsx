import React, { FC } from "react"
import { Case, Default, Switch } from "react-if"
import Article, { Kind } from "../../models/Article"
import ArticleContentTweet from "./ArticleContentTweet"
import ArticleContentSlideShare from "./ArticleContentSlideShare"
import ArticleContentYoutube from "./ArticleContentYoutube"
import ArticleContentMarkdown from "./ArticleContentMarkdown"


interface Props {
  article: Article
}

const ArticleContent: FC<Props> = ({article}) => (
  <Switch>
    <Case condition={article.kind === Kind.Tweet}>
      <ArticleContentTweet article={article}/>
    </Case>
    <Case condition={article.kind === Kind.SlideShare}>
      <ArticleContentSlideShare article={article}/>
    </Case>
    <Case condition={article.kind === Kind.Youtube}>
      <ArticleContentYoutube article={article}/>
    </Case>
    <Default>
      <ArticleContentMarkdown article={article}/>
    </Default>
  </Switch>
)

export default ArticleContent
