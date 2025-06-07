接口文档

## 统一响应格式：

- 正确响应：

```json
{
  "data": T,
  "requestId": string,
  "timeStamp": number
}
```

- 错误响应：

```json
{
  "code": number,
  "msg": string,
  "data": T,
  "requestId": string,
  "timeStamp": number
}
```

## account 账号模块

1. **registerAccount** 注册账号
   - 请求方式：POST
   - 请求路径：/api/v1/account/registerAccount
   - 请求参数 json：
     - email：string 类型，邮箱
     - nickname：string 类型，昵称
     - password：string 类型，密码
     - phone：string 类型，手机号
     - email_verification_code：string 类型，邮箱验证码
     - img_verification_code：string 类型，图片验证码，大小写不敏感
   - 响应示例：
     ```json
     {
       "data": {
         "nickname": "fender",
         "email": "927171598@qq.com"
       },
       "requestId": "TdGlsTqcsEBbUvhClaRQnAYXVbCRZjjB",
       "timeStamp": 1740052911
     }
     ```

2. **getAccount** 获取账号信息[须携带 token]
   - 请求方式：GET
   - 请求路径：/api/v1/account/getAccount?email=xxx
   - 请求参数 query：
     - email：string 类型，邮箱
   - 响应示例：
     ```json
     {
       "data": {
         "email": "927171598@qq.com",
         "nickname": "fender",
         "phone": "110"
       },
       "requestId": "FRjzgvEFXlsHWPKvvOCdtgAmiOidCwHt",
       "timeStamp": 1740053250
     }
     ```

3. **loginAccount** 登录账号
   - 请求方式：POST
   - 请求路径：/api/v1/account/loginAccount
   - 请求参数 json：
     - email：string 类型，邮箱
     - password：string 类型，密码
     - img_verification_code：string 类型，图片验证码，大小写不敏感
   - 响应示例：
    ```json
    {
      "data": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAwNjA3NTUsInVzZXJJZCI6Mn0.Ejv6v1ceDeArV-5zWjEExQUIwm-BfvwwHMRIH6hm3B4",
        "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDAyMjYzNTUsInVzZXJJZCI6Mn0.ZpbSXypjzG302IDff6BRLGM92Ieiywz8GimiZViwPr0"
      },
      "requestId": "WgXCIzQcTeEXXLFXKbxxTCrVVlnPsbvI",
      "timeStamp": 1740053555
    }
    ```

4. **logoutAccount** 登出账号[须携带 token]
   - 请求方式：POST
   - 请求路径：/api/v1/account/logoutAccount
   - 响应示例：
   ```json
   {
     "data": "用户注销成功",
     "requestId": "BNqxozHafYcfghrdbzaJyRMgZFiyUgee",
     "timeStamp": 1740063607
   }
   ```

5. **resetPassword** 重置密码[须携带 token]
   - 请求方式：POST
   - 请求路径：/api/v1/account/resetPassword
   - 请求参数 json：
     - email：string 类型，邮箱
     - new_password：string 类型，新密码
     - again_new_password：string 类型，再次输入新密码
     - email_verification_code：string 类型，邮箱验证码
   - 响应示例：
     ```json
     {
       "data": "密码重置成功",
       "requestId": "ZybJbcMxXCMJPhoJnZBABjiQMKTyvJNk",
       "timeStamp": 1740063697
     }
     ```

## post 文章模块

- 统一响应格式：

```json
{
  "data": {
    "id": number,
    "title": string,
    "image": string,
    "visibility": string,
    "content_html": string,
    "category_id": number,
    "gmt_create": string,
    "gmt_modified": string
  },
  "requestId": string,
  "timeStamp": number
}
```

> visibility 只有两种取值："public" 和 "private"，分别表示公开和私密。

1. **GetAllPosts** 获取包含所有文章的列表
   - 请求方式：GET
   - 请求路径：/api/v1/post/getAllPosts?page_size=xxx&page=xxx
   - 请求参数 query：
     - page_size：每页显示的文章数量
     - page：当前页码
   - 响应示例：
    ```json
    {
        "data": {
            "currentPage": 1,
            "posts": [
                {
                    "id": 1925164513678594048,
                    "title": "比特币如何挖矿（挖矿原理）-工作量证明",
                    "image": "https://haowallpaper.com/link/common/file/previewFileImg/15789130517090624",
                    "visibility": true,
                    "content_html": "<h1 id=\"-\">比特币如何挖矿（挖矿原理）-工作量证明</h1>\n<p>在<a href=\"https://learnblockchain.cn/2017/10/25/whatbc/\">区块链记账原理</a>一文中，我们了解到区块链记��",
                    "category_id": 1925162183231016960,
                    "gmt_create": "2025-05-26 19:06:32",
                    "gmt_modified": "2025-05-26 19:06:32"
                },
                {
                    "id": 1925164207242743808,
                    "title": "比特币所有权及隐私问题-非对称加密",
                    "image": "https://haowallpaper.com/link/common/file/previewFileImg/16893379459992960",
                    "visibility": true,
                    "content_html": "<h1 id=\"-\">比特币所有权及隐私问题-非对称加密</h1>\n<p>想象你有一个透明的保险箱，里面放着你的比特币。谁能打开这个箱子并取出里面的比特币？在传��",
                    "category_id": 1925162101823770624,
                    "gmt_create": "2025-05-26 19:06:32",
                    "gmt_modified": "2025-05-26 19:06:32"
                },
                {
                    "id": 1925163940988325888,
                    "title": "区块链记账原理",
                    "image": "https://haowallpaper.com/link/common/file/previewFileImg/16806298317868416",
                    "visibility": true,
                    "content_html": "<h1 id=\"heading\">区块链记账原理</h1>\n<p>想象一个魔法账本，每一页不仅记录交易，还与前一页用神奇墨水相连。这就是区块链记账的本质。</p>\n<p>区块链技�",
                    "category_id": 1925162101823770624,
                    "gmt_create": "2025-05-26 19:06:32",
                    "gmt_modified": "2025-05-26 19:06:32"
                },
                {
                    "id": 1925163668354371584,
                    "title": "什么是区块链",
                    "image": "https://haowallpaper.com/link/common/file/previewFileImg/16845070927449472",
                    "visibility": true,
                    "content_html": "<h1 id=\"heading\">什么是区块链？</h1>\n<p>想象一本永远无法篡改的数字账本，每一页都与前一页紧密相连，这就是区块链的基本概念。区块链是一种分布式数",
                    "category_id": 1925162101823770624,
                    "gmt_create": "2025-05-26 19:06:32",
                    "gmt_modified": "2025-05-26 19:06:32"
                },
                {
                    "id": 1925163154849927168,
                    "title": "接口文档",
                    "image": "https://haowallpaper.com/link/common/file/previewFileImg/16671743549427072",
                    "visibility": true,
                    "content_html": "<p>接口文档</p>\n<h2 id=\"heading\">统一响应格式：</h2>\n<ul>\n<li>正确响应：</li>\n</ul>\n<pre><code class=\"language-json\">{\n  &quot;data&quot;: T,\n  &quot;requestId&quot;: &quot;string&quot",
                    "category_id": 1925162017614729216,
                    "gmt_create": "2025-05-26 19:06:32",
                    "gmt_modified": "2025-05-26 19:06:32"
                }
            ],
            "totalPages": 2
        },
        "requestId": "JmsXSuYVYtBEcppJXZSuOBGlQOKDFzvP",
        "timeStamp": 1747832270
    }
    ```
   > 注：为了减少传输体积和提供预览效果，此接口对于 content_html 字段只会返回存储在数据库的 HTML 的前 200 个字符。

2. **getOnePost** 获取单篇文章详情
   - 请求方式：GET
   - 请求路径：/api/v1/post/getOnePost?id=xxx
   - 请求参数 query：
     - id：number 类型，文章 ID
   - 响应示例：
    ```json
    {
        "data": {
            "id": 1925164513678594048,
            "title": "比特币如何挖矿（挖矿原理）-工作量证明",
            "image": "https://haowallpaper.com/link/common/file/previewFileImg/15789130517090624",
            "visibility": true,
            "content_html": "<h1 id=\"-\">比特币如何挖矿（挖矿原理）-工作量证明</h1>\n<p>在<a href=\"https://learnblockchain.cn/2017/10/25/whatbc/\">区块链记账原理</a>一文中，我们了解到区块链记账是将交易记录、时间戳、账本序号和前一区块哈希值等信息打包计算哈希值的过程。这个过程需要消耗计算机资源，那么，为什么会有人愿意免费记账呢？</p>\n<p>答案很简单：<strong>谁帮大家记账，谁就能获得新铸造的比特币作为奖励</strong>。中本聪（比特币创始人）的这个设计很巧妙，既解决了记账问题，又实现了比特币的发行。这就像淘金热中的矿工们，付出劳动从河床中&quot;挖出&quot;黄金一样，因此这个过程被形象地称为&quot;挖矿&quot;。</p>\n<h2 id=\"heading\">记账工作：一场数学竞赛</h2>\n<p>既然记账有奖励，必然会有许多人争相参与。但如果多个人同时记账，账本就会出现不一致。比特币如何解决这个问题？方法是举办一场公平的&quot;数学竞赛&quot;：</p>\n<ul>\n<li>每 10 分钟左右只有一个&quot;幸运儿&quot;能获得记账权</li>\n<li>通过解决一道特别难的数学题（工作量证明）来决定谁是幸运儿</li>\n<li>其他人负责验证答案并接受获胜者的记账结果</li>\n</ul>\n<p>在参加这场竞赛前，矿工们需要做好以下准备工作：</p>\n<ul>\n<li>收集网络中尚未记录的有效交易</li>\n<li>检查每笔交易的付款地址是否有足够余额</li>\n<li>验证交易签名是否有效</li>\n<li>打包这些验证通过的交易</li>\n<li>添加一笔特殊交易：奖励自己一定数量的比特币</li>\n</ul>\n<p>如果成功获得记账权，这笔奖励就归矿工所有，就像挖到了一袋数字黄金。</p>\n<h2 id=\"heading-1\">工作量证明：寻找幸运数字的游戏</h2>\n<p>在区块链记账过程中，每次记账都会将前一区块的哈希值与当前账页信息一起计算哈希值。如果仅此而已，任何人都能轻松完成记账。</p>\n<p>为了控制记账速度（约 10 分钟一次），比特币系统设置了一道难题：<strong>找到一个特殊数字，使得整个区块的哈希值以特定数量的 0 开头</strong>。这就像要求抛硬币连续多次都是正面，概率非常小。为了寻找这个特殊数字，矿工需要引入一个可变参数——随机数（Nonce）。</p>\n<p>用伪代码表示就是：</p>\n<pre><code>Hash(前一区块哈希值, 交易记录集) = 456635BCD\n</code></pre>\n<p>加入随机数后：</p>\n<pre><code>Hash(前一区块哈希值, 交易记录集, 随机数) = 0000aFD635BCD\n</code></pre>\n<p>哈希函数有个特性：输入的微小变化会导致输出的巨大变化，就像一滴墨水能使整杯清水完全变色。矿工只能通过反复试验不同的随机数，期望碰巧找到一个满足条件的值。谁先找到，谁就赢得这轮记账权。</p>\n<h2 id=\"heading-2\">计算量分析：寻找宇宙级&quot;彩票号码&quot;</h2>\n<p>为什么说挖矿很难？让我们算算概率：</p>\n<p>哈希值包含数字和大小写字母，每个位置有 62 种可能（26 个大写字母+26 个小写字母+10 个数字）。如果要求第一位是 0，概率是 1/62，平均需要尝试 62 次才能碰上一次。</p>\n<p>如果要求前两位都是 0，概率变成 1/62²，需要尝试约 3844 次。要求前 n 位都是 0，则需尝试约 62ⁿ 次！这种难度呈指数级增长，就像要求掷骰子连续多次都掷出同一个点数。</p>\n<p>以区块#493050 为例：</p>\n<p><img src=\"https://img.learnblockchain.cn/2017/block_info_493050.jpg!wl\" alt=\"示例图片\" /></p>\n<p>数据来源：<a href=\"https://blockchain.info\">https://blockchain.info</a></p>\n<p>这个区块的哈希值以 18 个 0 开头！理论上需要尝试 62¹⁸ 次计算，这是一个天文数字——比宇宙中的原子总数还多。这就是为什么现在的矿工需要特殊设备和大量电力，而且通常会组成&quot;矿池&quot;共同挖矿，按贡献比例分配收益。</p>\n<p>从经济角度看，只要挖矿收益高于成本，就会有新矿工加入，加剧竞争，进一步提高难度。这就形成了一个自我调节的平衡系统。</p>\n<p>由于中国电力成本相对较低，中国矿工曾经控制了全网一半以上的算力。</p>\n<h2 id=\"heading-3\">验证：轻松确认的美妙设计</h2>\n<p>挖矿过程的一个绝妙之处在于：虽然找到解非常困难，但验证解是否正确却异常简单。</p>\n<p>当一个矿工找到符合条件的随机数后，会立即向全网广播新区块。其他矿工收到后，只需一次哈希计算就能验证其正确性，就像数独游戏——解题很费劲，但检查答案却很容易。</p>\n<p>验证通过后，其他矿工会接受这个新区块，将它加入自己的账本，然后立即转向竞争下一个区块的记账权。这确保了全网账本的一致性。</p>\n<p>如果有人试图作弊（例如记录虚假交易），其区块会在验证环节被拒绝，导致前期投入的大量计算资源白白浪费。这种经济惩罚机制使得遵守规则比作弊更有利可图，保障了整个系统的安全。</p>\n<p>关于区块结构如何验证交易的详细信息，可参考<a href=\"https://xiaozhuanlan.com/topic/1402935768\">比特币区块结构 Merkle 树及简单支付验证分析</a>。</p>\n<h2 id=\"heading-4\">说明</h2>\n<p>矿工的收入不仅包括新发行的比特币奖励，还包括用户支付的交易费。这些奖励机制共同激励矿工维护网络安全。</p>\n<p>图中红箭头标示的是本文所涉及的信息。</p>\n<p>比特币共识协议主要由工作量证明和最长链机制两部分组成，详情请参考<a href=\"https://xiaozhuanlan.com/topic/0298513746\">比特币如何达成共识 - 最长链的选择</a>。</p>\n",
            "category_id": 1925162183231016960,
            "gmt_create": "2025-05-26 19:06:32",
            "gmt_modified": "2025-05-26 19:06:32"
        },
        "requestId": "HrFmQJGBsTpkjXdwKJIFINtRbwimPxTr",
        "timeStamp": 1747832314
    }
    ```
   > 注：此接口对于 content_html 字段会返回完整的 HTML 内容。

3. **createOnePost** 创建文章[须携带 token]
   - 请求方式：POST
   - 请求路径：/api/v1/post/createOnePost
   - 请求参数 form-data：
     - title：string 类型，文章标题
     - image：string 类型，文章图片 URL
     - visibility：boolean 类型，文章可见性，取值：0 或 1，也可以 false 或 true，0 表示私密，1 表示公开
     - content_markdown: string 类型，文章内容的 Markdown 格式
     - category_id：number 类型，文章所属类目 ID
   - 响应示例：
     ```json
     {
       "data": {
         "id": 1917209986715357184,
         "title": "测试文档002",
         "image": "https://haowallpaper.com/link/common/file/previewFileImg/16445196366630272",
         "visibility": true,
         "content_html": "<h1 id=\"golang1\">Golang小项目(1)</h1>\n<hr />\n<h2 id=\"heading\">前言</h2>\n<blockquote>\n<p>本项目适合Golang初学者，通过简单的项目实践来加深对Golang的基本语法和Web开发的理解。</p>\n</blockquote>\n<h2 id=\"heading-1\">正文</h2>\n<h3 id=\"heading-2\">项目结构</h3>\n<pre><code class=\"language-Golang\">.\r\n├── main.go\r\n└── static\r\n    ├── form.html\r\n    └── index.html\r\n</code></pre>\n<h3 id=\"heading-3\">项目流程图</h3>\n<p><img src=\"https://s2.loli.net/2024/09/23/Oy8uaXmoMPd9NV7.png\" alt=\"\" /></p>\n<blockquote>\n<p>定义三个路由：</p>\n</blockquote>\n<ul>\n<li><code>/</code>：首页，显示<code>static/index.html</code>页面</li>\n<li><code>/hello</code>：欢迎页面，显示 <code>hello</code></li>\n<li><code>/form</code>：表单页面，显示<code>static/form.html</code>页面</li>\n</ul>\n<h3 id=\"heading-4\">项目初始化</h3>\n<pre><code class=\"language-Golang\">package main\r\n\r\n// 导入必要的包\r\nimport (\r\n    &quot;fmt&quot;\r\n    &quot;log&quot;\r\n    &quot;net/http&quot;\r\n)\r\n\r\nfunc main(){\r\n\r\n}\r\n</code></pre>\n<blockquote>\n<p>在static目录下创建index.html</p>\n</blockquote>\n<pre><code class=\"language-html\">&lt;!DOCTYPE html&gt;\r\n&lt;html lang=&quot;en&quot;&gt;\r\n&lt;head&gt;\r\n    &lt;meta charset=&quot;UTF-8&quot;&gt;\r\n    &lt;meta name=&quot;viewport&quot; content=&quot;width=device-width, initial-scale=1.0&quot;&gt;\r\n    &lt;title&gt;Static Website&lt;/title&gt;\r\n&lt;/head&gt;\r\n&lt;body&gt;\r\n    &lt;h2&gt;Static Website&lt;/h2&gt;\r\n&lt;/body&gt;\r\n&lt;/html&gt;\r\n</code></pre>\n<blockquote>\n<p>在static目录下创建form.html</p>\n</blockquote>\n<pre><code class=\"language-html\">&lt;!DOCTYPE html&gt;\r\n&lt;html lang=&quot;en&quot;&gt;\r\n&lt;head&gt;\r\n    &lt;meta charset=&quot;UTF-8&quot;&gt;\r\n    &lt;meta name=&quot;viewport&quot; content=&quot;width=device-width, initial-scale=1.0&quot;&gt;\r\n    &lt;title&gt;Form Page&lt;/title&gt;\r\n&lt;/head&gt;\r\n&lt;body&gt;\r\n    &lt;div&gt;\r\n        &lt;form action=&quot;/form&quot; method=&quot;POST&quot;&gt;\r\n            &lt;label&gt;Name&lt;/label&gt;&lt;input name=&quot;name&quot; type=&quot;text&quot; value=&quot;&quot;/&gt;\r\n            &lt;label&gt;Address&lt;/label&gt;&lt;input name=&quot;address&quot; type=&quot;text&quot; value=&quot;&quot;/&gt;\r\n\r\n            &lt;input type=&quot;submit&quot; value=&quot;Submit&quot;/&gt;\r\n        &lt;/form&gt;\r\n    &lt;/div&gt;\r\n&lt;/body&gt;\r\n&lt;/html&gt;\r\n</code></pre>\n<h3 id=\"heading-5\">编写程序主入口</h3>\n<pre><code class=\"language-Golang\">func main() {\r\n    // 创建一个 HTTP 文件服务器，用于提供静态文件\r\n    fileServer := http.FileServer(http.Dir(&quot;./static&quot;))\r\n    // 将根路径 / 绑定到 fileServer 处理器。这意味着所有根路径的请求都由 fileServer 处理，提供 ./static 目录中的文件。\r\n    http.Handle(&quot;/&quot;, fileServer)\r\n    // 将 /form 路径绑定到 formHandler 函数。这意味着对 /form 的请求将由 formHandler 处理。\r\n\thttp.HandleFunc(&quot;/form&quot;, formHandler)\r\n    // 将 /hello 路径绑定到 helloHandler 函数。这意味着对 /hello 的请求将由 helloHandler 处理。\r\n    http.HandleFunc(&quot;/hello&quot;, helloHandler)\r\n\r\n    // 打印一条消息，表示服务器正在启动，监听 8080 端口。\r\n    fmt.Fprintf(&quot;Starting server at port 8080\\n&quot;)\r\n    // 启动一个 HTTP 服务器，监听 8080 端口。如果启动失败，err 将不为空。\r\n    // nil 表示使用默认的 ServeMux 路由器。\r\n    if err := http.ListenAndServe(&quot;:8080&quot;,nil); err != nil {\r\n        // 如果服务器启动失败，记录错误并终止程序。\r\n        log.Fatal(err)\r\n    }\r\n}\r\n</code></pre>\n<h3 id=\"heading-6\">编写路由处理函数</h3>\n<pre><code class=\"language-Golang\">// 表单处理函数\r\n// w 为 http.ResponseWriter 的实例，用于向客户端返回响应。\r\n// r 为 http.Request 的实例，包含了客户端的请求信息。\r\nfunc formHandler(w http.ResponseWriter, r *http.Request){\r\n\t// 尝试解析表单数据。如果解析失败，将打印err信息。\r\n    if err := r.ParseForm(); err != nil {\r\n\t\tfmt.Fprintf(w, &quot;ParseForm() err: %v&quot;, err)\r\n        // 结束函数执行\r\n\t\treturn\r\n\t}\r\n    // 如果表单解析成功，向响应写入成功信息。\r\n\tfmt.Fprintf(w, &quot;POST request successful&quot;)\r\n    // 获取form.html页面中表单中名为 name 和 address 的字段值。\r\n\tname := r.FormValue(&quot;name&quot;)\r\n\taddress := r.FormValue(&quot;address&quot;)\r\n    // 将 name 和 address 字段值写入响应 w。\r\n\tfmt.Fprintf(w, &quot;Name = %s\\n&quot;, name)\r\n\tfmt.Fprintf(w, &quot;Address = %s\\n&quot;, address)\r\n}\r\n\r\n\r\n\r\n/// 欢迎页面处理函数\r\nfunc helloHandler(w http.ResponseWriter, r *http.Request){\r\n    // 检查request请求的 URL 路径是否为 /hello，如果不是，返回 404 错误响应。\r\n\tif r.URL.Path != &quot;/hello&quot; {\r\n\t\thttp.Error(w, &quot;404 not found&quot;, http.StatusNotFound)\r\n        // 结束函数执行\r\n\t\treturn\r\n\t}\r\n    // 检查请求的方法是否为 GET，如果不是，返回 405 错误（即请求方法不允许）。\r\n\tif r.Method != &quot;GET&quot; {\r\n\t\thttp.Error(w, &quot;method is not supported&quot;, http.StatusNotFound)\r\n\t\treturn\r\n\t}\r\n    //  如果路径和方法均匹配，向响应写入 &quot;hello!&quot;。\r\n\tfmt.Fprintf(w, &quot;hello!&quot;)\r\n}\r\n</code></pre>\n<h3 id=\"heading-7\">运行项目</h3>\n<pre><code class=\"language-Golang\">go build main.go\r\ngo run main.go\r\n</code></pre>\n<blockquote>\n<p>打开浏览器，访问 http://localhost:8080/ ，显示 <code>Static Website</code> 页面。<br />\n打开浏览器，访问 http://localhost:8080/hello ，显示 <code>hello!</code> 页面。<br />\n打开浏览器，访问 http://localhost:8080/form.html ，显示 <code>Form Page</code> 页面，并可以输入姓名和地址。</p>\n</blockquote>\n<h2 id=\"heading-8\">每日小技巧：指针传递</h2>\n<blockquote>\n<p>在 Go 语言中，* 用于声明和操作指针。具体到 r *http.Request 中的 *，它表示 r 是一个指向 http.Request 类型的指针。下面详细解释了在这种上下文中使用 * 的含义和作用：</p>\n<ol>\n<li>声明指针</li>\n</ol>\n<blockquote>\n<p>在函数参数中使用 * 来声明指针类型。这意味着参数 r 是一个 http.Request 的指针，而不是 http.Request 的值本身。</p>\n<pre><code class=\"language-Golang\">func handleRequest(r *http.Request) {\r\n    // 这里的 r 是 *http.Request 类型\r\n}\r\n</code></pre>\n</blockquote>\n<ol start=\"2\">\n<li>访问和操作 http.Request 对象</li>\n</ol>\n<blockquote>\n<ul>\n<li><strong>通过指针访问字段</strong>: 使用指针 *http.Request 可以直接访问和修改 http.Request 对象的字段，因为指针指向了实际的 http.Request 对象，而不是其副本。</li>\n<li><strong>修改请求</strong>: 如果函数需要对请求进行修改（例如修改请求头或解析请求体），使用指针可以直接在原始对象上进行操作。</li>\n</ul>\n<pre><code class=\"language-Golang\">func modifyHeader(r *http.Request) {\r\n    r.Header.Set(&quot;X-Custom-Header&quot;, &quot;Value&quot;) // 修改请求头\r\n}\r\n</code></pre>\n</blockquote>\n<ol start=\"3\">\n<li>避免对象复制</li>\n</ol>\n<blockquote>\n<ul>\n<li><strong>节省内存</strong>: 如果 http.Request 是一个大结构体，直接传递指针而不是整个结构体可以减少内存使用和复制开销。</li>\n<li><strong>提高性能</strong>: 传递指针避免了将整个结构体复制到函数调用栈，从而提高了性能。</li>\n</ul>\n</blockquote>\n<ol start=\"4\">\n<li>指针解引用</li>\n</ol>\n<blockquote>\n<ul>\n<li><strong>获取实际值</strong>: 如果你有一个指向 http.Request 的指针，可以通过解引用（*r）来访问实际的 http.Request 对象。</li>\n</ul>\n<pre><code class=\"language-Golang\">func printRequestMethod(r *http.Request) {\r\nfmt.Println(r.Method) // 访问请求方法\r\n}\r\n</code></pre>\n</blockquote>\n<ol start=\"5\">\n<li>一致性和标准库</li>\n</ol>\n<blockquote>\n<ul>\n<li><strong>与标准库一致</strong>: Go 的标准库中处理 HTTP 请求的函数普遍使用 *http.Request，这为开发者提供了一种一致的编程模式，并确保了对 http.Request 的高效处理。</li>\n</ul>\n</blockquote>\n</blockquote>\n",
         "category_id": 0,
         "gmt_create": "2025-05-26 19:06:32",
         "gmt_modified": "2025-05-26 19:06:32"
       },
       "requestId": "eyvAxQQXOiBumMRGWwFjNDXKLRNYMTzI",
       "timeStamp": 1745933455
     }
     ```

4. **updateOnePost** 更新文章[须携带 token]
   - 请求方式：POST
   - 请求路径：/api/v1/post/updateOnePost
   - 请求参数 json：
     - id：number 类型，文章 ID
     - title：string 类型，文章标题
     - image：string 类型，文章图片 URL
     - visibility：boolean 类型，文章可见性，取值：0 或 1，也可以 false 或 true，0 表示私密，1 表示公开
     - content_markdown: string 类型，文章内容的 Markdown 格式，支持文件路径和直接输入 markdown 文件内容
     - category_id：number 类型，文章所属类目 ID
       > 除了 id 为必填项外，其他字段都为可选，只会更新传递的字段，未传递的字段保持原值。
   - 响应示例：
       ```json
       {
         "data": {
           "id": 1917209260593254400,
           "title": "测试文档更新成功",
           "image": "https://haowallpaper.com/link/common/file/previewFileImg/16445196366630272",
           "visibility": true,
           "content_html": "<h1 id=\"golang1\">Golang小项目(1)</h1>\n<hr />\n<h2 id=\"heading\">前言</h2>\n<blockquote>\n<p>本项目适合Golang初学者，通过简单的项目实践来加深对Golang的基本语法和Web开发的理解。</p>\n</blockquote>\n<h2 id=\"heading-1\">正文</h2>\n<h3 id=\"heading-2\">项目结构</h3>\n<pre><code class=\"language-Golang\">.\r\n├── main.go\r\n└── static\r\n    ├── form.html\r\n    └── index.html\r\n</code></pre>\n<h3 id=\"heading-3\">项目流程图</h3>\n<p><img src=\"https://s2.loli.net/2024/09/23/Oy8uaXmoMPd9NV7.png\" alt=\"\" /></p>\n<blockquote>\n<p>定义三个路由：</p>\n</blockquote>\n<ul>\n<li><code>/</code>：首页，显示<code>static/index.html</code>页面</li>\n<li><code>/hello</code>：欢迎页面，显示 <code>hello</code></li>\n<li><code>/form</code>：表单页面，显示<code>static/form.html</code>页面</li>\n</ul>\n<h3 id=\"heading-4\">项目初始化</h3>\n<pre><code class=\"language-Golang\">package main\r\n\r\n// 导入必要的包\r\nimport (\r\n    &quot;fmt&quot;\r\n    &quot;log&quot;\r\n    &quot;net/http&quot;\r\n)\r\n\r\nfunc main(){\r\n\r\n}\r\n</code></pre>\n<blockquote>\n<p>在static目录下创建index.html</p>\n</blockquote>\n<pre><code class=\"language-html\">&lt;!DOCTYPE html&gt;\r\n&lt;html lang=&quot;en&quot;&gt;\r\n&lt;head&gt;\r\n    &lt;meta charset=&quot;UTF-8&quot;&gt;\r\n    &lt;meta name=&quot;viewport&quot; content=&quot;width=device-width, initial-scale=1.0&quot;&gt;\r\n    &lt;title&gt;Static Website&lt;/title&gt;\r\n&lt;/head&gt;\r\n&lt;body&gt;\r\n    &lt;h2&gt;Static Website&lt;/h2&gt;\r\n&lt;/body&gt;\r\n&lt;/html&gt;\r\n</code></pre>\n<blockquote>\n<p>在static目录下创建form.html</p>\n</blockquote>\n<pre><code class=\"language-html\">&lt;!DOCTYPE html&gt;\r\n&lt;html lang=&quot;en&quot;&gt;\r\n&lt;head&gt;\r\n    &lt;meta charset=&quot;UTF-8&quot;&gt;\r\n    &lt;meta name=&quot;viewport&quot; content=&quot;width=device-width, initial-scale=1.0&quot;&gt;\r\n    &lt;title&gt;Form Page&lt;/title&gt;\r\n&lt;/head&gt;\r\n&lt;body&gt;\r\n    &lt;div&gt;\r\n        &lt;form action=&quot;/form&quot; method=&quot;POST&quot;&gt;\r\n            &lt;label&gt;Name&lt;/label&gt;&lt;input name=&quot;name&quot; type=&quot;text&quot; value=&quot;&quot;/&gt;\r\n            &lt;label&gt;Address&lt;/label&gt;&lt;input name=&quot;address&quot; type=&quot;text&quot; value=&quot;&quot;/&gt;\r\n\r\n            &lt;input type=&quot;submit&quot; value=&quot;Submit&quot;/&gt;\r\n        &lt;/form&gt;\r\n    &lt;/div&gt;\r\n&lt;/body&gt;\r\n&lt;/html&gt;\r\n</code></pre>\n<h3 id=\"heading-5\">编写程序主入口</h3>\n<pre><code class=\"language-Golang\">func main() {\r\n    // 创建一个 HTTP 文件服务器，用于提供静态文件\r\n    fileServer := http.FileServer(http.Dir(&quot;./static&quot;))\r\n    // 将根路径 / 绑定到 fileServer 处理器。这意味着所有根路径的请求都由 fileServer 处理，提供 ./static 目录中的文件。\r\n    http.Handle(&quot;/&quot;, fileServer)\r\n    // 将 /form 路径绑定到 formHandler 函数。这意味着对 /form 的请求将由 formHandler 处理。\r\n\thttp.HandleFunc(&quot;/form&quot;, formHandler)\r\n    // 将 /hello 路径绑定到 helloHandler 函数。这意味着对 /hello 的请求将由 helloHandler 处理。\r\n    http.HandleFunc(&quot;/hello&quot;, helloHandler)\r\n\r\n    // 打印一条消息，表示服务器正在启动，监听 8080 端口。\r\n    fmt.Fprintf(&quot;Starting server at port 8080\\n&quot;)\r\n    // 启动一个 HTTP 服务器，监听 8080 端口。如果启动失败，err 将不为空。\r\n    // nil 表示使用默认的 ServeMux 路由器。\r\n    if err := http.ListenAndServe(&quot;:8080&quot;,nil); err != nil {\r\n        // 如果服务器启动失败，记录错误并终止程序。\r\n        log.Fatal(err)\r\n    }\r\n}\r\n</code></pre>\n<h3 id=\"heading-6\">编写路由处理函数</h3>\n<pre><code class=\"language-Golang\">// 表单处理函数\r\n// w 为 http.ResponseWriter 的实例，用于向客户端返回响应。\r\n// r 为 http.Request 的实例，包含了客户端的请求信息。\r\nfunc formHandler(w http.ResponseWriter, r *http.Request){\r\n\t// 尝试解析表单数据。如果解析失败，将打印err信息。\r\n    if err := r.ParseForm(); err != nil {\r\n\t\tfmt.Fprintf(w, &quot;ParseForm() err: %v&quot;, err)\r\n        // 结束函数执行\r\n\t\treturn\r\n\t}\r\n    // 如果表单解析成功，向响应写入成功信息。\r\n\tfmt.Fprintf(w, &quot;POST request successful&quot;)\r\n    // 获取form.html页面中表单中名为 name 和 address 的字段值。\r\n\tname := r.FormValue(&quot;name&quot;)\r\n\taddress := r.FormValue(&quot;address&quot;)\r\n    // 将 name 和 address 字段值写入响应 w。\r\n\tfmt.Fprintf(w, &quot;Name = %s\\n&quot;, name)\r\n\tfmt.Fprintf(w, &quot;Address = %s\\n&quot;, address)\r\n}\r\n\r\n\r\n\r\n/// 欢迎页面处理函数\r\nfunc helloHandler(w http.ResponseWriter, r *http.Request){\r\n    // 检查request请求的 URL 路径是否为 /hello，如果不是，返回 404 错误响应。\r\n\tif r.URL.Path != &quot;/hello&quot; {\r\n\t\thttp.Error(w, &quot;404 not found&quot;, http.StatusNotFound)\r\n        // 结束函数执行\r\n\t\treturn\r\n\t}\r\n    // 检查请求的方法是否为 GET，如果不是，返回 405 错误（即请求方法不允许）。\r\n\tif r.Method != &quot;GET&quot; {\r\n\t\thttp.Error(w, &quot;method is not supported&quot;, http.StatusNotFound)\r\n\t\treturn\r\n\t}\r\n    //  如果路径和方法均匹配，向响应写入 &quot;hello!&quot;。\r\n\tfmt.Fprintf(w, &quot;hello!&quot;)\r\n}\r\n</code></pre>\n<h3 id=\"heading-7\">运行项目</h3>\n<pre><code class=\"language-Golang\">go build main.go\r\ngo run main.go\r\n</code></pre>\n<blockquote>\n<p>打开浏览器，访问 http://localhost:8080/ ，显示 <code>Static Website</code> 页面。<br />\n打开浏览器，访问 http://localhost:8080/hello ，显示 <code>hello!</code> 页面。<br />\n打开浏览器，访问 http://localhost:8080/form.html ，显示 <code>Form Page</code> 页面，并可以输入姓名和地址。</p>\n</blockquote>\n<h2 id=\"heading-8\">每日小技巧：指针传递</h2>\n<blockquote>\n<p>在 Go 语言中，* 用于声明和操作指针。具体到 r *http.Request 中的 *，它表示 r 是一个指向 http.Request 类型的指针。下面详细解释了在这种上下文中使用 * 的含义和作用：</p>\n<ol>\n<li>声明指针</li>\n</ol>\n<blockquote>\n<p>在函数参数中使用 * 来声明指针类型。这意味着参数 r 是一个 http.Request 的指针，而不是 http.Request 的值本身。</p>\n<pre><code class=\"language-Golang\">func handleRequest(r *http.Request) {\r\n    // 这里的 r 是 *http.Request 类型\r\n}\r\n</code></pre>\n</blockquote>\n<ol start=\"2\">\n<li>访问和操作 http.Request 对象</li>\n</ol>\n<blockquote>\n<ul>\n<li><strong>通过指针访问字段</strong>: 使用指针 *http.Request 可以直接访问和修改 http.Request 对象的字段，因为指针指向了实际的 http.Request 对象，而不是其副本。</li>\n<li><strong>修改请求</strong>: 如果函数需要对请求进行修改（例如修改请求头或解析请求体），使用指针可以直接在原始对象上进行操作。</li>\n</ul>\n<pre><code class=\"language-Golang\">func modifyHeader(r *http.Request) {\r\n    r.Header.Set(&quot;X-Custom-Header&quot;, &quot;Value&quot;) // 修改请求头\r\n}\r\n</code></pre>\n</blockquote>\n<ol start=\"3\">\n<li>避免对象复制</li>\n</ol>\n<blockquote>\n<ul>\n<li><strong>节省内存</strong>: 如果 http.Request 是一个大结构体，直接传递指针而不是整个结构体可以减少内存使用和复制开销。</li>\n<li><strong>提高性能</strong>: 传递指针避免了将整个结构体复制到函数调用栈，从而提高了性能。</li>\n</ul>\n</blockquote>\n<ol start=\"4\">\n<li>指针解引用</li>\n</ol>\n<blockquote>\n<ul>\n<li><strong>获取实际值</strong>: 如果你有一个指向 http.Request 的指针，可以通过解引用（*r）来访问实际的 http.Request 对象。</li>\n</ul>\n<pre><code class=\"language-Golang\">func printRequestMethod(r *http.Request) {\r\nfmt.Println(r.Method) // 访问请求方法\r\n}\r\n</code></pre>\n</blockquote>\n<ol start=\"5\">\n<li>一致性和标准库</li>\n</ol>\n<blockquote>\n<ul>\n<li><strong>与标准库一致</strong>: Go 的标准库中处理 HTTP 请求的函数普遍使用 *http.Request，这为开发者提供了一种一致的编程模式，并确保了对 http.Request 的高效处理。</li>\n</ul>\n</blockquote>\n</blockquote>\n",
           "category_id": 0,
           "gmt_create": "2025-05-26 19:06:32",
           "gmt_modified": "2025-05-26 19:06:32"
         },
         "requestId": "WOPPPRpOGjEeuVayRvNybbPMiHkrJbXi",
         "timeStamp": 1745933757
       }
       ```

5. **deleteOnePost** 删除文章[须携带 token]
   - 请求方式：POST
   - 请求路径：/api/v1/post/deleteOnePost
   - 请求参数 json：
     - id：number 类型，文章 ID
   - 响应示例：
     ```json
     {
       "data": "文章删除成功",
       "requestId": "zWaMCAOkBoYiojZppBSJYZDDNnkCCmrr",
       "timeStamp": 1740048955
     }
     ```

## category 类目模块

- 统一响应格式：

```json
{
  "data": {
    "id": number,
    "name": string,
    "description": string,
    "parent_id": number,
    "path": string,
    "children": number
  },
  "requestId": string,
  "timeStamp": number
}
```

1. **getOneCategory** 获取单个类目详情
   - 请求方式：GET
   - 请求路径：/api/v1/category/getOneCategory?id=xxx
   - 请求参数 query：
     - id：number 类型，类目 ID
   - 响应示例：
    ```json
    {
      "data": {
        "id": 1925162017614729216,
        "name": "测试类目001",
        "description": "测试类目001",
        "parent_id": 0,
        "path": "",
        "children": null
      },
      "requestId": "rAlMXMaKchFehlupdkQzGkqbwhPwjwPr",
      "timeStamp": 1747831528
    }
    ```

2. **getCategoryTree** 获取类目树
   - 请求方式：GET
   - 请求路径：/api/v1/category/getCategoryTree
   - 响应示例：
    ```json
    {
      "data": [
        {
          "id": 1925162017614729216,
          "name": "测试类目001",
          "description": "测试类目001",
          "parent_id": 0,
          "path": "",
          "children": [
            {
              "id": 1925162101823770624,
              "name": "测试类目002",
              "description": "测试类目002",
              "parent_id": 1925162017614729216,
              "path": "/1925162017614729216",
              "children": [
                {
                  "id": 1925162183231016960,
                  "name": "测试类目003",
                  "description": "测试类目003",
                  "parent_id": 1925162101823770624,
                  "path": "/1925162017614729216/1925162101823770624",
                  "children": []
                }
              ]
            }
          ]
        },
        {
          "id": 1925162244161671168,
          "name": "测试类目004",
          "description": "测试类目004",
          "parent_id": 0,
          "path": "",
          "children": []
        },
        {
          "id": 1925162320862908416,
          "name": "测试类目005",
          "description": "测试类目005",
          "parent_id": 0,
          "path": "",
          "children": []
        }
      ],
      "requestId": "QRDPHofmSJUEnGhjzWBasuxgtSFjcVgY",
      "timeStamp": 1747831453
    }
    ```

3. **createOneCategory** 创建类目[须携带 token]
   - 请求方式：POST
   - 请求路径：/api/v1/category/createOneCategory
   - 请求参数 json：
     - name：string 类型，类目名称
     - description：string 类型，类目描述
     - parent_id：number 类型，父类目 ID
   - 响应示例：
    ```json
    {
      "data": {
        "id": 1925171262762520576,
        "name": "测试类目007",
        "description": "测试类目007",
        "parent_id": 1925171198778413056,
        "path": "/1925171198778413056",
        "children": null
      },
      "requestId": "bnHBGyDMIichmYQlUtNvuINIDgNgRIHP",
      "timeStamp": 1747831571
    }
    ```

4. **updateOneCategory** 更新类目[须携带 token]
   - 请求方式：POST
   - 请求路径：/api/v1/category/updateOneCategory
   - 请求参数 json：
     - id：number 类型，类目 ID
     - name：string 类型，类目名称
     - description：string 类型，类目描述
     - parent_id：number 类型，父类目 ID，根类目为 0，不传则不修改父类目
   - 响应示例：
    ```json
    {
      "data": {
        "id": 1925171262762520576,
        "name": "测试类目007-改",
        "description": "测试类目007-改",
        "parent_id": 1925171198778413056,
        "path": "/1925171198778413056",
        "children": null
      },
      "requestId": "bnHBGyDMIichmYQlUtNvuINIDgNgRIHP",
      "timeStamp": 1747831571
    }
    ```
   > 注：更新类目时，如果不传递 parent_id 字段，则表示不修改父类目，如果 parent_id 字段传 0，则表示将类目置于根类目下。

5. **deleteOneCategory** 删除类目[须携带 token]
   - 请求方式：POST
   - 请求路径：/api/v1/category/deleteOneCategory
   - 请求参数 json：
     - id：number 类型，类目 ID
   - 响应示例：
    ```json
    {
      "data": [
        {
          "id": 1925171262762520576,
          "name": "测试类目007",
          "description": "测试类目007",
          "parent_id": 1925171198778413056,
          "path": "/1925171198778413056",
          "children": []
        }
      ],
      "requestId": "ZaieRRfEKxrmbTIyedVvQBdUhKfvhnBo",
      "timeStamp": 1747831660
    }
    ```

6. **getCategoryChildrenTree** 获取类目子树
   - 请求方式：GET
   - 请求路径：/api/v1/category/getCategoryChildrenTree?id=xxx
   - 请求参数 query：
     - id：number 类型，类目 ID
   - 响应示例：
    ```json
    {
      "data": [
        {
          "id": 1925162101823770624,
          "name": "测试类目002",
          "description": "测试类目002",
          "parent_id": 1925162017614729216,
          "path": "/1925162017614729216",
          "children": [
            {
              "id": 1925162183231016960,
              "name": "测试类目003",
              "description": "测试类目003",
              "parent_id": 1925162101823770624,
              "path": "/1925162017614729216/1925162101823770624",
              "children": []
            }
          ]
        }
      ],
      "requestId": "klvwKJVTEztovlGosHOJFEBmyWuVjGNb",
      "timeStamp": 1747831716
    }
    ```

## verification 验证码模块

1. **SendImgVerificationCode** 发送图形验证码
   - 请求方式：GET
   - 请求路径：/api/v1/verification/sendImgVerificationCode?email=xxx
   - 请求参数 query：
     - email：string 类型，邮箱地址
   - 响应示例：
     ```json
     {
       "data": {
         "imgBase64": "data:image/png;base64,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"
       },
       "requestId": "kXmCrFwABkpNyMjsvhmSrQmyZeXPdhrh",
       "timeStamp": 1740049437
     }
     ```

2. **SendEmailVerificationCode** 验证邮箱验证码
   - 请求方式：GET
   - 请求路径：/api/v1/verification/sendEmailVerificationCode?email=xxx
   - 请求参数 query：
     - email：string 类型，邮箱地址
   - 响应示例：
     ```json
     {
       "data": "邮箱验证码发送成功, 请注意查收！",
       "requestId": "tPFcJhDOJSHXDSMdULLfRlGDHFMShFYe",
       "timeStamp": 1740049546
     }
     ```

## comment 评论模块

1. **getOneComment** 获取单条评论
   - 请求方式：GET
   - 请求路径：/api/v1/comment/getOneComment?id=xxx
   - 请求参数 query：
     - id：string 类型，评论 ID
   - 响应示例：
    ```json
    {
      "data": {
        "id": 1925165213221392384,
        "content": "测试评论001",
        "user_id": 1,
        "post_id": 1925162665324318720,
        "reply_to_comment_id": 0,
        "replies": null
      },
      "requestId": "FXknjuOOElRyCJckPnOILXJOCAcQdnpb",
      "timeStamp": 1747831980
    }
    ```

2. **getCommentGraph** 获取文章下的评论
   - 请求方式：GET
   - 请求路径：/api/v1/comment/getCommentGraph?post_id=xxx
   - 请求参数 query：
     - post_id：string 类型，文章 ID
   - 响应示例：
    ```json
    {
      "data": [
        {
          "id": 1925165213221392384,
          "content": "测试评论001",
          "user_id": 1,
          "post_id": 1925162665324318720,
          "reply_to_comment_id": 0,
          "replies": [
            {
              "id": 1925165308734083072,
              "content": "测试评论002",
              "user_id": 1,
              "post_id": 1925162665324318720,
              "reply_to_comment_id": 1925165213221392384,
              "replies": [
                {
                  "id": 1925165374127476736,
                  "content": "测试评论003",
                  "user_id": 1,
                  "post_id": 1925162665324318720,
                  "reply_to_comment_id": 1925165308734083072,
                  "replies": []
                }
              ]
            }
          ]
        },
        {
          "id": 1925165448005947392,
          "content": "测试评论004",
          "user_id": 1,
          "post_id": 1925162665324318720,
          "reply_to_comment_id": 0,
          "replies": []
        }
      ],
      "requestId": "NQJOarmmEeZAFmWyTYdyFlsafxzXTFtp",
      "timeStamp": 1747831853
    }
    ```

3. **createOneComment** 创建评论
   - 请求方式：POST
   - 请求路径：/api/v1/comment/createOneComment
   - 请求参数 json：
     - content：string 类型，评论内容
     - user_id：number 类型，用户 ID
     - post_id：number 类型，文章 ID
     - reply_to_comment_id：number 类型，回复的评论 ID
       > 注：reply_to_comment_id 为 0 时，表示对文章进行评论，reply_to_comment_id 不为 0 时，表示对文章的评论进行回复，默认为 0
   - 响应示例：
    ```json
    {
      "data": {
        "id": 1925165213221392384,
        "content": "测试评论001",
        "user_id": 1,
        "post_id": 1925162665324318720,
        "reply_to_comment_id": 0,
        "replies": null
      },
      "requestId": "FXknjuOOElRyCJckPnOILXJOCAcQdnpb",
      "timeStamp": 1747831980
    }
    ```

4. **deleteOneComment** 删除评论
   - 请求方式：POST
   - 请求路径：/api/v1/comment/deleteOneComment
   - 请求参数 json：
     - id：number 类型，评论 ID
   - 响应示例：
    ```json
    {
      "data": {
        "id": 1925165213221392384,
        "content": "测试评论001",
        "user_id": 1,
        "post_id": 1925162665324318720,
        "reply_to_comment_id": 0,
        "replies": null
      },
      "requestId": "FXknjuOOElRyCJckPnOILXJOCAcQdnpb",
      "timeStamp": 1747831980
    }
    ```

## oss 模块

1. **uploadOneFile** 上传文件[须携带 token]

   - 请求方式：POST
   - 请求路径：/api/v1/oss/uploadOneFile
   - 请求参数 form-data：
     - bucket_name：string 类型，桶名称
     - upload_file：string 类型，上传文件
   - 响应示例：

   ```json
   {
     "data": {
       "object_path": "/jank/开发者文档_1924859223817064448.md"
     },
     "requestId": "jUbgtOWGPAAkBaADzCBNiSLIdrXWGPmY",
     "timeStamp": 1747757175
   }
   ```

2. **downloadOneFile** 下载文件[须携带 token]

   - 请求方式：GET
   - 请求路径：/api/v1/oss/downloadOneFile?bucket_name=xxx&object_name=xxx
   - 请求参数 query：
     - bucket_name：string 类型，桶名称
     - object_name：string 类型，对象名称
   - 响应示例：

   ```json
   {
     "data": {
       "file_name": "开发者文档_1924839933462188032.md",
       "download_url": "http://127.0.0.1:9001/jank/%E5%BC%80%E5%8F%91%E8%80%85%E6%96%87%E6%A1%A3_1924839933462188032.md?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=2EGruziH9sx4L7zItUHU%2F20250520%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20250520T160643Z&X-Amz-Expires=172800&X-Amz-SignedHeaders=host&response-content-disposition=attachment%3B%20filename%2A%3DUTF-8%27%27%25E5%25BC%2580%25E5%258F%2591%25E8%2580%2585%25E6%2596%2587%25E6%25A1%25A3_1924839933462188032.md&response-content-encoding=UTF-8&response-content-type=text%2Fmarkdown&X-Amz-Signature=4199e260c06bfcdc378c3c59d943db38dcc70a22d7c45e872c616a77a62dffc5",
       "expires_at": "2025-05-23 00:06:43"
     },
     "requestId": "sufuLYtODJVlnFCevaFpZFrnnuTHMHll",
     "timeStamp": 1747757203
   }
   ```

3. **deleteOneFile** 删除文件[须携带 token]

   - 请求方式：POST
   - 请求路径：/api/v1/oss/deleteOneFile
   - 请求参数 json：
     - bucket_name：string 类型，桶名称
     - object_name：string 类型，对象名称
   - 响应示例：

   ```json
   {
     "data": "文件删除成功",
     "requestId": "MNTUrtBGHTaKAlaqKTJdFHqxUPVSTbAW",
     "timeStamp": 1747757378
   }
   ```

4. **listAllObjects** 获取指定桶内的所有对象[须携带 token]
   - 请求方式：GET
   - 请求路径：/api/v1/oss/listAllObjects?bucket_name=xxx&prefix=xxx
   - 请求参数 query：
     - bucket_name：string 类型，桶名称
     - prefix：string 类型，对象名称
     > 其中 prefix 为可选参数，如果不指定则列出所有对象。
   - 响应示例：
   ```json
   {
     "data": {
       "objects": [
         "开发者文档_1924839933462188032.md",
         "开发者文档_1924859223817064448.md"
       ]
     },
     "requestId": "lpffIokeyetcfeqzSesTsecMScKPthjt",
     "timeStamp": 1747757539
   }
   ```

## test 测试模块

1. **testPing** 测试接口
   - 请求方式：GET
   - 请求路径：/api/v1/test/testPing
   - 请求参数：无
   - 响应示例：
   ```text
   Pong successfully!
   ```
2. **testHello** 测试接口
   - 请求方式：GET
   - 请求路径：/api/v1/test/testHello
   - 请求参数：无
   - 响应示例：
   ```text
   Hello, Jank 🎉!
   ```
3. **testLogger** 测试接口
   - 请求方式：GET
   - 请求路径：/api/v1/test/testLogger
   - 请求参数：无
   - 响应示例：
   ```text
   测试日志成功!
   ```
4. **testRedis** 测试接口
   - 请求方式：GET
   - 请求路径：/api/v1/test/testRedis
   - 请求参数：无
   - 响应示例：
   ```text
   测试缓存功能完成!
   ```
5. **testSuccessRes** 测试接口
   - 请求方式：GET
   - 请求路径：/api/v1/test/testSuccessRes
   - 请求参数：无
   - 响应示例：
   ```json
   {
     "data": "测试成功响应成功!",
     "requestId": "XtZvqFlDtpgzwEAesJpFMGgJQRbQDXyM",
     "timeStamp": 1740118491
   }
   ```
6. **testErrRes** 测试接口
   - 请求方式：GET
   - 请求路径：/api/v1/test/testErrRes
   - 请求参数：无
   - 响应示例：
   ```json
   {
     "code": 10000,
     "msg": "服务端异常",
     "data": {},
     "requestId": "BRnzCMxAoprBllAuBGPWqoDNofArbuOX",
     "timeStamp": 1740118534
   }
   ```
7. **testErrorMiddleware** 测试接口

   - 请求方式：GET
   - 请求路径：/api/v1/test/testErrorMiddleware
   - 请求参数：无
   - 响应示例：无

8. **testLongReq** 测试接口
   - 请求方式：GET
   - 请求路径：/api/v2/test/testLongReq
   - 请求参数：无
   - 响应示例：
   ```text
   模拟耗时请求处理完成!
   ```
