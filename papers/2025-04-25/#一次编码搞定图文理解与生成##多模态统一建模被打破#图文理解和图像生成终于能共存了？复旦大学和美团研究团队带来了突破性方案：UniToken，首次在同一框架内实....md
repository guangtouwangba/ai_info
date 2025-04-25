# #一次编码搞定图文理解与生成##多模态统一建模被打破#图文理解和图像生成终于能共存了？复旦大学和美团研究团队带来了突破性方案：UniToken，首次在同一框架内实...

**URL**: https://weibo.com/6105753431/PoRX7DqAm

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E4%B8%80%E6%AC%A1%E7%BC%96%E7%A0%81%E6%90%9E%E5%AE%9A%E5%9B%BE%E6%96%87%E7%90%86%E8%A7%A3%E4%B8%8E%E7%94%9F%E6%88%90%23&amp;extparam=%23%E4%B8%80%E6%AC%A1%E7%BC%96%E7%A0%81%E6%90%9E%E5%AE%9A%E5%9B%BE%E6%96%87%E7%90%86%E8%A7%A3%E4%B8%8E%E7%94%9F%E6%88%90%23" data-hide=""><span class="surl-text">#一次编码搞定图文理解与生成#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%9A%E6%A8%A1%E6%80%81%E7%BB%9F%E4%B8%80%E5%BB%BA%E6%A8%A1%E8%A2%AB%E6%89%93%E7%A0%B4%23&amp;extparam=%23%E5%A4%9A%E6%A8%A1%E6%80%81%E7%BB%9F%E4%B8%80%E5%BB%BA%E6%A8%A1%E8%A2%AB%E6%89%93%E7%A0%B4%23" data-hide=""><span class="surl-text">#多模态统一建模被打破#</span></a><br><br>图文理解和图像生成终于能共存了？<br><br>复旦大学和美团研究团队带来了突破性方案：UniToken，首次在同一框架内实现理解与生成“双优表现”的统一视觉编码器，直接打破多模态统一建模的老大难。<br><br>核心创新点如下：<br><br>- 连续+离散双编码器：既保留高层语义（适合理解），又能还原底层细节（适合生成）。<br>- 统一格式：用一种结构同时表达图像的两种表示，再加文本，模型可自由切换任务。<br>- 三阶段训练流程：先对齐视觉语义空间，再联合理解/生成训练，最后微调提升复杂指令跟随能力。<br><br>实验表现亮眼：<br><br>- 在多个图文理解和图像生成的权威评测中，UniToken不但没拉胯，反而表现出色，甚至超越不少单一任务专精模型。<br>- 数据量越大、任务越复杂，UniToken展现的鲁棒性越强。<br><br>代码和模型全都开放，点击文章查看： <a href="https://weibo.com/ttarticle/p/show?id=2309405159272697954485" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">UniToken：多模态AI的“全能选手”，一次编码搞定图文理解与图像生成</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0suwpvrtkj30rs0fmq7m.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

复旦大学与美团团队提出UniToken模型，突破性地实现了多模态统一建模。该方案通过连续+离散双编码器结构，既能保留高层语义（适合图文理解），又能还原底层细节（适合图像生成）。采用三阶段训练流程，在多个权威评测中同时超越单一任务模型，且数据量越大表现越优。研究成果已开源，标志着AI在统一处理图文理解与生成任务上取得重要进展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-25T03:17:17Z
- **目录日期**: 2025-04-25
