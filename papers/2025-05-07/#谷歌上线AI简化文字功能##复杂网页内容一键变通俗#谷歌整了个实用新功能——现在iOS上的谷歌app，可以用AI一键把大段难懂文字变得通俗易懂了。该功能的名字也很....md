# #谷歌上线AI简化文字功能##复杂网页内容一键变通俗#谷歌整了个实用新功能——现在iOS上的谷歌app，可以用AI一键把大段难懂文字变得通俗易懂了。该功能的名字也很...

**URL**: https://weibo.com/6105753431/PqHxBs9ib

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%B0%B7%E6%AD%8C%E4%B8%8A%E7%BA%BFAI%E7%AE%80%E5%8C%96%E6%96%87%E5%AD%97%E5%8A%9F%E8%83%BD%23&amp;extparam=%23%E8%B0%B7%E6%AD%8C%E4%B8%8A%E7%BA%BFAI%E7%AE%80%E5%8C%96%E6%96%87%E5%AD%97%E5%8A%9F%E8%83%BD%23" data-hide=""><span class="surl-text">#谷歌上线AI简化文字功能#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%8D%E6%9D%82%E7%BD%91%E9%A1%B5%E5%86%85%E5%AE%B9%E4%B8%80%E9%94%AE%E5%8F%98%E9%80%9A%E4%BF%97%23&amp;extparam=%23%E5%A4%8D%E6%9D%82%E7%BD%91%E9%A1%B5%E5%86%85%E5%AE%B9%E4%B8%80%E9%94%AE%E5%8F%98%E9%80%9A%E4%BF%97%23" data-hide=""><span class="surl-text">#复杂网页内容一键变通俗#</span></a><br><br>谷歌整了个实用新功能——<br><br>现在iOS上的谷歌app，可以用AI一键把大段难懂文字变得通俗易懂了。该功能的名字也很直白，就叫“Simplify（简化）”。<br><br>它依赖的是谷歌自家的Gemini模型，让你在浏览网页时，无需跳转，就能读懂晦涩内容。<br><br>比如医学报告里看到“emphysema”和“fibrosis”这类术语，Simplify会在不改变原意的情况下，给出语言更亲民、更易懂的解释。【图1】<br><br>那么，它是怎么做到的？技术上可分为三步：【图2】<br><br>- 先让Gemini生成简化版本；  <br>- 再用另一个Gemini模型从可读性、完整性、符合原意三个维度打分；  <br>- 分数反馈后再优化prompt，重复这个循环直到找到效果最好的写法。<br><br>整个过程跑了800多轮迭代，才得出现在你看到的这个简化效果。<br><br>目前这个功能已经在iOS版谷歌app上线，用法很简单：网页上选中你觉得难懂的文字，点击弹出来的Simplify按钮，就能看到简化版内容。<br><br>谷歌没有透露是否会扩展到Android或桌面版Chrome，但表示会“持续考虑推广到更多产品中”。<br><br>技术报告：<a href="https://research.google/blog/making-complex-text-understandable-minimally-lossy-text-simplification-with-gemini/" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span> <span class="surl-text">网页链接</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i16sv2bnszg30rs0gnb2m.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i16ss3hrjpj30es0gltcr.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

谷歌在iOS版应用中推出"Simplify"AI功能，基于Gemini模型可将复杂文本(如医学术语)简化为通俗表达。技术实现分三步：Gemini生成简化文本→另一模型从可读性/完整性/准确性三方面评分→优化提示词循环迭代800多次。用户选中网页文字即可一键转换，目前仅限iOS端，未来可能扩展至其他平台。该功能无需跳转页面，直接在原网页提供易懂解释。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-07T06:02:54Z
- **目录日期**: 2025-05-07
