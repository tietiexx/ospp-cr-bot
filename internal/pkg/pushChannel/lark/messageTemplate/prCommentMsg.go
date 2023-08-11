package messageTemplate

const PrCommentTitle1 = "有人在 Comment 中 @ 你!"
const PrCommentTitle2 = "有 Comment 待处理!"
const PrCommentMsg = `{
  "config": {
    "wide_screen_mode": true
  },
  "header": {
    "template": "blue",
    "title": {
      "content": "⏰ {{Title}}",
      "tag": "plain_text"
    }
  },
  "elements": [
    {
      "fields": [
        {
          "is_short": true,
          "text": {
            "content": "**🕐 时间：**\n{{Time}}",
            "tag": "lark_md"
          }
        },
        {
          "is_short": true,
          "text": {
            "content": "**📋 PR 标题：**\n{{PRTitle}}",
            "tag": "lark_md"
          }
        },
        {
          "is_short": false,
          "text": {
            "content": "",
            "tag": "lark_md"
          }
        },
        {
          "is_short": true,
          "text": {
            "content": "**👤 提出人：**\n{{Login}}",
            "tag": "lark_md"
          }
        },
        {
          "is_short": true,
          "text": {
            "content": "**🔢 PR 号：**\n{{PRNumber}}",
            "tag": "lark_md"
          }
        },
        {
          "is_short": false,
          "text": {
            "content": "",
            "tag": "lark_md"
          }
        },
        {
          "is_short": true,
          "text": {
            "content": "**🔢 提醒人：**\n{{RequestedReviewers}}",
            "tag": "lark_md"
          }
        }
      ],
      "tag": "div"
    },
    {
      "tag": "markdown",
      "content": "**📜 PR 链接：**\n[]({{PRContent}})"
    },
    {
      "tag": "hr"
    },
    {
      "tag": "action",
      "actions": [
        {
          "tag": "button",
          "text": {
            "tag": "plain_text",
            "content": "跟进 PR"
          },
          "url": "{{PRURL}}",
          "type": "primary"
        }
      ]
    }
  ]
}`
