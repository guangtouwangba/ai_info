# #资深工程师的LLM技巧##专业工程师如何使用LLM#专业的工程师都怎么使用LLM？BuzzFeed的资深数据科学家Max Woolf坦率地在博客里分享了自己的经验，量子位进行了简...

**URL**: https://weibo.com/6105753431/PqJuWpcA4

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%B5%84%E6%B7%B1%E5%B7%A5%E7%A8%8B%E5%B8%88%E7%9A%84LLM%E6%8A%80%E5%B7%A7%23&amp;extparam=%23%E8%B5%84%E6%B7%B1%E5%B7%A5%E7%A8%8B%E5%B8%88%E7%9A%84LLM%E6%8A%80%E5%B7%A7%23" data-hide=""><span class="surl-text">#资深工程师的LLM技巧#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%93%E4%B8%9A%E5%B7%A5%E7%A8%8B%E5%B8%88%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8LLM%23&amp;extparam=%23%E4%B8%93%E4%B8%9A%E5%B7%A5%E7%A8%8B%E5%B8%88%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8LLM%23" data-hide=""><span class="surl-text">#专业工程师如何使用LLM#</span></a><br><br>专业的工程师都怎么使用LLM？<br><br>BuzzFeed的资深数据科学家Max Woolf坦率地在博客里分享了自己的经验，量子位进行了简短的总结，一起来看看吧👇<br><br>一、 Woolf如何使用LLM？<br><br>- 利用提示工程进行结果优化<br>精细设计提示词以提升输出质量，提示工程已成为LLM使用者的必备技能。<br><br>- 直接使用API接口<br>避免使用面向普通用户的前端界面，直接调用各LLM服务商提供的后端UI，这些轻量级封装接口既能保留API的全部功能，又便于必要时移植到代码中。<br>像ChatGPT API这类原生接口，可以设置系统提示词来精细控制生成规则，通过API还能控制生成的温度参数，调节输出的随机性。<br><br>二、LLM解决问题的优势和不足<br><br>LLM能够完成80%的基础工作，但剩余的20%（包括迭代优化、测试验证和收集反馈）反而耗时更长。即便模型输出日趋稳定，“幻觉问题”仍是隐忧。<br><br>三、在什么情况下不使用LLM？<br>- 写作<br>LLM无法模仿博客的个人风格，最重要的是，科技/编程领域的最新动态很难出现在LLM训练数据中，这会增加幻觉风险。<br>不过，写作方面倒是有一个技巧：将文章喂给LLM，要求它扮演尖刻的评论员对文章内容进行评论。它的反馈能暴露论点薄弱处，迫使自己思考解决方案。<br><br>- 陪伴<br>一个被训练得极尽友善、却因幻觉问题习惯性撒谎的实体，实在很难成为真正的朋友。<br><br>四、如何在写代码时利用LLM？<br>- LLM在编程时的应用场景<br>1、高效解决简单问题：在生成正则表达式等简单任务上，LLM非常高效<br>2、提供更量身定制的解答：对于一些在Stack Overflow找不到现成答案的问题，LLM能提供更加详细的结果。<br>3、复杂问题的初步思路：至于更复杂的问题，LLM可以提供初步的代码框架和思路，虽然生成的代码需要进一步测试和改进，但提供了有用的起点和创新的想法。<br><br>- LLM编程会遇到的问题<br>1、对特定库支持不足：LLM在处理pandas和polars库时，可能会混淆函数，导致需要查阅文档来确认。<br>2、数学运算结果不可靠：LLM无法准确输出数学运算的结果，而利用代码解释器来处理大规模数据，成本过于高昂。<br>3、编程助手会分散注意力：Copilot弹出的代码建议会打乱编程时的专注状态。<br><br>- 对智能体、MCP和氛围编码的看法<br>1、智能体与MCP：如今的智能体工作流确实更可靠了，MCP标准化也比传统工具范式更有优势，但说到底，它们并没有创造出什么新场景，无非是把两年前LangChain刚问世时就能实现的功能，用更复杂晦涩的方式重新实现了一遍。<br><br>2、氛围编码：氛围编码或许能完成80%的工作，但若以“氛围编码”为借口，在重要项目中故意交付劣质代码，这种行为就极不专业。<br><br>五、如何看待LLM的未来<br><br>LLM只是一个工具，要做的是在适当的时候使用正确的工具。<br><br>受限于篇幅，内容有所删减，想阅读更多朋友们欢迎点击：<br><a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fminimaxir.com%2F2025%2F05%2Fllm-use%2F" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i171h1p5pgj30l00l07c0.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i171h3dvhaj30xn0zkqgg.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

资深工程师Max Woolf分享使用LLM的核心技巧：1) 通过提示工程优化输出，直接调用API接口控制参数；2) LLM擅长解决80%基础工作（如生成正则表达式、提供代码框架），但存在幻觉问题，剩余20%需人工优化；3) 避免用于个人风格写作/陪伴，但可作批判性审阅；4) 编程时注意特定库支持不足、数学运算不可靠等问题；5) 智能体工具未突破性创新。他强调LLM是辅助工具，需合理运用。关键点：API调用、80/20法则、警惕幻觉、工具定位。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T18:03:28Z
- **目录日期**: 2025-05-07
