# #AI画图前能先思考了##港中文让AI文生图也能推理#AI绘画，也开始“动脑子”了!香港中文大学MMLab团队推出了T2I-R1，这是首个基于强化学习、具备推理能力的文生图...

**URL**: https://weibo.com/6105753431/PrDFfvZB6

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E7%94%BB%E5%9B%BE%E5%89%8D%E8%83%BD%E5%85%88%E6%80%9D%E8%80%83%E4%BA%86%23&amp;extparam=%23AI%E7%94%BB%E5%9B%BE%E5%89%8D%E8%83%BD%E5%85%88%E6%80%9D%E8%80%83%E4%BA%86%23" data-hide=""><span class="surl-text">#AI画图前能先思考了#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%B8%AF%E4%B8%AD%E6%96%87%E8%AE%A9AI%E6%96%87%E7%94%9F%E5%9B%BE%E4%B9%9F%E8%83%BD%E6%8E%A8%E7%90%86%23&amp;extparam=%23%E6%B8%AF%E4%B8%AD%E6%96%87%E8%AE%A9AI%E6%96%87%E7%94%9F%E5%9B%BE%E4%B9%9F%E8%83%BD%E6%8E%A8%E7%90%86%23" data-hide=""><span class="surl-text">#港中文让AI文生图也能推理#</span></a><br><br>AI绘画，也开始“动脑子”了!<br><br>香港中文大学MMLab团队推出了T2I-R1，这是首个基于强化学习、具备推理能力的文生图模型。<br><br>T2I-R1的核心创新是“双层推理框架”：<br><br>- 语义层推理在生成图像前规划整体结构，比如物体外观和位置；<br><br>- Token级推理则处理图像生成过程中的细节和连贯性。<br><br>配合这个双层框架，T2I-R1引入了BiCoT-GRPO策略，一种新型强化学习方法，能在一次训练中同步优化两层推理过程，提高效率、降低成本。<br><br>另一个难点是图像生成的质量评估。T2I-R1用多个视觉专家模型组成奖励机制，从多个维度评估图像质量，既确保稳定性，也防止模型过拟合。<br><br>实验证明，T2I-R1在多个文生图评测中全面超越已有模型，甚至在一些任务上打败了最强基线FLUX.1。更重要的是，它生成的图像更贴近人类预期，对复杂场景也更鲁棒。详情请看文章： <a href="https://weibo.com/ttarticle/p/show?id=2309405165873299325010" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">文生图进入R1时代：港中文MMLab发布T2I-R1，让AI绘画“先推理再下笔</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1dxcwps1lj30rr0fmac6.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

香港中文大学MMLab团队开发出首个具备推理能力的文生图模型T2I-R1。该模型采用"双层推理框架"：语义层预先规划图像整体结构，Token级处理生成细节。配合新型强化学习方法BiCoT-GRPO同步优化两层推理，并引入多专家模型评估生成质量。实验显示T2I-R1在多项评测中超越现有模型，生成的图像更符合人类预期，对复杂场景表现更稳健。这项技术使AI绘画实现了"先思考再创作"的突破。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T09:03:11Z
- **目录日期**: 2025-05-13
