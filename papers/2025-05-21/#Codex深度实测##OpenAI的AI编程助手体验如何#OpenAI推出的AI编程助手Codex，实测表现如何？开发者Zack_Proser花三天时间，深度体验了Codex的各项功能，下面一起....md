# #Codex深度实测##OpenAI的AI编程助手体验如何#OpenAI推出的AI编程助手Codex，实测表现如何？开发者Zack Proser花三天时间，深度体验了Codex的各项功能，下面一起...

**URL**: https://weibo.com/6105753431/PsPo2vvcv

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Codex%E6%B7%B1%E5%BA%A6%E5%AE%9E%E6%B5%8B%23&amp;extparam=%23Codex%E6%B7%B1%E5%BA%A6%E5%AE%9E%E6%B5%8B%23" data-hide=""><span class="surl-text">#Codex深度实测#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23OpenAI%E7%9A%84AI%E7%BC%96%E7%A8%8B%E5%8A%A9%E6%89%8B%E4%BD%93%E9%AA%8C%E5%A6%82%E4%BD%95%23&amp;extparam=%23OpenAI%E7%9A%84AI%E7%BC%96%E7%A8%8B%E5%8A%A9%E6%89%8B%E4%BD%93%E9%AA%8C%E5%A6%82%E4%BD%95%23" data-hide=""><span class="surl-text">#OpenAI的AI编程助手体验如何#</span></a><br><br>OpenAI推出的AI编程助手Codex，实测表现如何？<br><br>开发者Zack Proser花三天时间，深度体验了Codex的各项功能，下面一起来看看。<br><br>首先是主页面，映入眼帘的，是一个大大的任务输入框。下方按顺序排列着各项任务，还支持任务归档（Archive），右上角有着Environments（环境）和官方Docs（文档）。【图1】<br><br>在任务列表内，你可以看到任务名，灰色字体还标出了任务时间和项目路径，右侧则显示了任务当前的状态。【图1】<br><br>在输入框输入需求后，用户可以在下方选项卡，选择具体的任务以及目标分支（main），即团队协作过程中，多个开发人员修改的记录。【图2】<br><br>再来看手机页面。【图3】<br><br>CodeX在保留核心功能的基础上，对页面做了极简适配，真正做到随时随地，掏出手机就能写代码，还能跑的那种。【图3】<br><br>点进一个具体的任务，页面顶部显示了此前对话记录，紧接着则是任务耗时（5m9s），下方则汇报了具体的执行情况，如增加搜索按钮、优化产品样式、插入评论区等。【图4】<br><br>除了对话框，上方选项卡还支持切换Diff（代码改动）和Logs（日志）。<br><br>点进Diff，主页面中，红色部分是删除代码，绿色部分是新增代码，上方显示了当前文件名和变动统计（新增+10，删除-6）。主页面左侧，用户可以继续与CodeX对话。【图5】<br><br>编辑好后，点击右上角Push，会弹出四个选项：【图6】<br><br>- Create New PR：创建正式的Pull Request，发起合并请求。<br>- Create PR (draft)：创建草稿状态的PR，适合需要进一步修改的提交。<br>- Copy git apply：复制Git补丁应用命令，供本地开发者手动应用。<br>- Copy patch：复制原始补丁内容（patch 文件），也可以手动粘贴应用。<br><br>虽然整体流程很清晰，但测试者也吐槽了当前的不足：<br><br>- 复杂任务能力有限：一旦任务涉及多轮逻辑、分支协同或大型重构，Codex就容易卡壳。<br>- 测试环境受限：Codex当前运行在“无网络沙盒”中，像npm install、curl等操作都跑不了，遇到依赖问题，只能人工补刀。<br>- 错误处理反馈不清晰：有时候任务失败，但UI不会提示具体问题，日志需要自己翻。<br>- 无法持续在一个分支修改：多轮更新会生成多个PR，不适合“我想一步步完善这个feature”的开发流程，只适合一次性完成小任务。<br><br>整体来看，Codex体验可谓相当惊艳，但细节尚需打磨。<br><br>原文：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fzackproser.com%2Fblog%2Fopenai-codex-review" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1mywyfyntj31fa1cwail.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1myx02c2cj31yi12ak05.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1myx36pupj30u01uoqhe.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1myx4pzt3j30u01uo156.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1myx61n5oj327u1cw4qp.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1myx6wu10j315i16ugsq.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

OpenAI的Codex编程助手实测显示其界面简洁，支持任务管理和代码修改可视化（红删绿增），提供PR创建、补丁复制等功能，手机端适配良好。但存在复杂任务处理能力有限、测试环境受限（无法执行网络操作）、错误反馈不明确、无法持续分支修改等问题。整体体验惊艳但需优化细节。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T08:04:35Z
- **目录日期**: 2025-05-21
