# #Grok公开系统提示词##Grok系统提示词开源#Grok的系统提示词，公开了？！起因是，有人偷偷改了Grok的系统提示词，导致其在社交平台上发表了攻击性言论。为了挽回...

**URL**: https://weibo.com/6105753431/Ps45dF9lN

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Grok%E5%85%AC%E5%BC%80%E7%B3%BB%E7%BB%9F%E6%8F%90%E7%A4%BA%E8%AF%8D%23&amp;extparam=%23Grok%E5%85%AC%E5%BC%80%E7%B3%BB%E7%BB%9F%E6%8F%90%E7%A4%BA%E8%AF%8D%23" data-hide=""><span class="surl-text">#Grok公开系统提示词#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Grok%E7%B3%BB%E7%BB%9F%E6%8F%90%E7%A4%BA%E8%AF%8D%E5%BC%80%E6%BA%90%23&amp;extparam=%23Grok%E7%B3%BB%E7%BB%9F%E6%8F%90%E7%A4%BA%E8%AF%8D%E5%BC%80%E6%BA%90%23" data-hide=""><span class="surl-text">#Grok系统提示词开源#</span></a><br><br>Grok的系统提示词，公开了？！<br><br>起因是，有人偷偷改了Grok的系统提示词，导致其在社交平台上发表了攻击性言论。<br><br>为了挽回信任，xAI官方彻查了事件源头，并宣布从现在开始，Grok的所有提示词完全开源，大家可以随时在GitHub上查看每一次变动。<br><br>而且官方还加了24小时监控机制和审核流程，防止再被人“偷偷刀”。<br><br>官方公开的文件共四个：<br><br>- ask_grok_summarizer.j2：这是用于Grok总结X帖子内容时的提示词，语气被设计成“极度怀疑主义”，鼓励多角度看问题，同时还有限制，比如“不能引用原帖”、“最多750字符”、“不要自称或下判断”。<br><br>- default_deepsearch_final_summarizer_prompt.j2：这个用于DeepSearch功能，也就是Grok在联网搜索后进行总结时的提示词。里面强调，搜索结果只是初步信息，不能直接当作事实，输出要有分析判断。<br><br>- grok3_official0330_p1.j2：这是聊天模型的系统提示词，相当于Grok的大脑指令集，定义了它的语气风格、回答逻辑，比如“保持中立”、“用段落组织信息”“不要盲目相信搜索结果”等等。<br>   <br>- grok_analyze_button.j2：这个是X平台“Explain”按钮背后的提示词，当用户点击时，Grok会用这个模板来生成解释。内容偏简明直白，适合公开场合快速说明。<br><br>值得注意的是，这些提示词是用Jinja2模板写的，能根据变量动态调整，比如插入当前时间、用户问题、分析结果等。<img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1h63fmza5j314c168kd4.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1h63gn4k9j31eo1cye1n.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

xAI官方公开了Grok AI的系统提示词源代码，起因是有人篡改提示词导致不当言论。为提升透明度，xAI在GitHub开源了四个核心提示词模板，涵盖内容总结、联网搜索分析、聊天逻辑和"Explain"按钮功能。这些Jinja2模板强调中立性、多角度分析和事实核查，并新增24小时监控防止篡改。公开文件显示Grok被设计为保持怀疑态度，限制主观判断，且对搜索结果持审慎态度。此举旨在重建信任，展示AI决策过程。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-16T05:03:21Z
- **目录日期**: 2025-05-16
