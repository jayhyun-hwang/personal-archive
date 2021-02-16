import React, { FC } from "react"
import { Button } from "rsuite"
import { useHistory } from "react-router-dom"
import MarkdownContent from "../../component/common/MarkdownContent"
import Article from "../../models/Article"
import GithubRepoWidget from "../../component/GithubRepoWidget"


interface Props {
  article: Article
}

const ArticleContentMarkdown: FC<Props> = ({article}) => {
  const history = useHistory()

  return (
    <>
      <div style={{textAlign: 'right'}}>
        <Button
          appearance="link"
          onClick={() => history.push(`/articles/${article.id}/edit`)}
        >
          EDIT
        </Button>
      </div>
      {renderWidget(article.url)}
      <MarkdownContent content={article.content}/>
    </>
  )
}

const renderWidget = (url: string) => {
  const githubRepoRegex = /^https:\/\/github\.com\/([a-zA-Z0-9-]*)\/([a-zA-Z0-9-]*)/gi
  const matched = githubRepoRegex.exec(url) || []
  if (matched.length < 3) {
    return null
  }

  const [ user, repo ] = [ matched[1], matched[2] ]

  return <GithubRepoWidget user={user} repo={repo} />
}

export default ArticleContentMarkdown
