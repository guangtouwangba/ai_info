# #一图总结五大智能体工作流##图解五大智能体工作流#智能体都已经不陌生，关于它的工作流，你了解多少呢？去年，Anthropic曾对这些智能体系统的工作流进行了总结...

**URL**: https://weibo.com/6105753431/PqJJy5eL9

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E5%9B%BE%E6%80%BB%E7%BB%93%E4%BA%94%E5%A4%A7%E6%99%BA%E8%83%BD%E4%BD%93%E5%B7%A5%E4%BD%9C%E6%B5%81%23&amp;extparam=%23%E4%B8%80%E5%9B%BE%E6%80%BB%E7%BB%93%E4%BA%94%E5%A4%A7%E6%99%BA%E8%83%BD%E4%BD%93%E5%B7%A5%E4%BD%9C%E6%B5%81%23" data-hide=""><span class="surl-text">#一图总结五大智能体工作流#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%9B%BE%E8%A7%A3%E4%BA%94%E5%A4%A7%E6%99%BA%E8%83%BD%E4%BD%93%E5%B7%A5%E4%BD%9C%E6%B5%81%23&amp;extparam=%23%E5%9B%BE%E8%A7%A3%E4%BA%94%E5%A4%A7%E6%99%BA%E8%83%BD%E4%BD%93%E5%B7%A5%E4%BD%9C%E6%B5%81%23" data-hide=""><span class="surl-text">#图解五大智能体工作流#</span></a><br><br>智能体都已经不陌生，关于它的工作流，你了解多少呢？<br><br>去年，Anthropic曾对这些智能体系统的工作流进行了总结。Aurimas Griciūna将关键内容整理在了一张图上，我们一起复习一下👇<br><br>一、提示链 (Prompt Chaining)<br>这种模式将复杂的任务分解成一系列步骤，并通过将它们串联起来解决问题，一个LLM调用的输出成为另一个LLM调用的输入。在大多数情况下，提示链以适度增加延迟为代价换取更高准确率。<br>适用场景：能够清晰拆分为固定子任务的场景。<br>示例：编写文档大纲，并核验大纲是否符合特定标准基于，最终大纲完成文档撰写。<br><br>二、路由 (Routing)<br>在这种模式中，输入会被分类，并被定向到专门的后续任务中。<br>适用场景：当工作流较为复杂，且特定拓扑路径通过专用工作流能更高效处理时，该方案尤为适用。若缺乏该机制，针对某类输入的优化反而可能影响其他输入的处理性能。<br>示例：将不同类型的客户服务查询（一般问题、退款请求、技术支持）引导到不同的下游流程、提示模版和处理工具中。<br><br>三、并行化 (Parallelization)<br>初始输入被拆分为多个查询，传递给LLM，然后聚合答案以产生最终答案。<br>适用场景：当速度很重要，并且被分解的子任务可以并行处理时非常有用。此外，当需要多个视角或多次尝试来获得更高置信度的结果时，并行化是有效的。<br>示例：从发票中提取多个项目，所有项目都可以并行进一步处理以提高速度。<br><br>四、编排器 (Orchestrator)<br>一个中央LLM动态地分解任务并委派给其他LLM或子工作流。<br>适用场景：无法预知需要哪些子任务的复杂场景。<br>示例：涉及从多个来源收集和分析可能相关信息的搜索任务。<br><br>五、评估器-优化器 (Evaluator-optimizer)<br>生成器LLM产生结果，然后评估器LLM对其进行评估，并在必要时提供反馈以进行进一步改进。<br>适用场景：对于需要持续改进的任务非常有用。<br>示例：文学翻译中存在着一些细微之处，翻译的LLM可能最初无法捕捉到，但评估的LLM可以提供有用的批评。<br><br>想查看更多内容？欢迎重温Anthropic去年年底发布的详细文章：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.anthropic.com%2Fengineering%2Fbuilding-effective-agents" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1725phffzj30v00zkb2a.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Anthropic总结了五大智能体工作流：1）提示链（Prompt Chaining）：将任务分解为串联步骤，提高准确性；2）路由（Routing）：分类输入并定向到专用流程；3）并行化（Parallelization）：拆分任务并行处理以提升速度；4）编排器（Orchestrator）：动态分解任务并委派子流程；5）评估器-优化器（Evaluator-optimizer）：生成结果后评估反馈以优化。这些模式适用于不同场景，如复杂任务拆分、多路径处理或持续改进需求。详情可参考Anthropic原文。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T04:04:55Z
- **目录日期**: 2025-05-08
