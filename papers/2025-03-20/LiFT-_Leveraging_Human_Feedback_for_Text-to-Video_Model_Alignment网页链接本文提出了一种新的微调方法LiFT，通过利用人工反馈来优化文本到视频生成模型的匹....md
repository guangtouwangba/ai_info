# LiFT: Leveraging Human Feedback for Text-to-Video Model Alignment网页链接本文提出了一种新的微调方法LiFT，通过利用人工反馈来优化文本到视频生成模型的匹...

**URL**: https://weibo.com/1870858943/P4kWe0L4G

## 原始摘要

LiFT: Leveraging Human Feedback for Text-to-Video Model Alignment<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.aminer.cn%2Fpub%2F675659efae8580e7ff8d68a8%2F%3Ff%3Dwb" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>本文提出了一种新的微调方法LiFT，通过利用人工反馈来优化文本到视频生成模型的匹配度。研究团队首先构建了一个包含大约1万个由人工评分及其理由组成的人类评分注释数据集LiFT-HRA。基于此数据集，训练了一个奖励模型LiFT-Critic，该模型可以有效地学习奖励函数，作为人类判断的代理，衡量给定视频与人类期望之间的匹配度。最后，利用学到的奖励函数通过最大化奖励加权的似然性来调整T2V模型。以CogVideoX-2B为案例，研究结果显示经过微调的模型在所有16项指标上均优于CogVideoX-5B，证明了人类反馈在提高生成视频的匹配度和质量方面的潜力。<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%9F%E6%88%90%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#生成模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A5%96%E5%8A%B1%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E5%A5%96%E5%8A%B1%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#奖励模型#</span></a><a href="https://m.weibo.cn/p/index?extparam=%E4%BA%BA%E5%B7%A5%E6%99%BA%E8%83%BD&amp;containerid=100808f068f0dad74789bee210163c40a4b50d" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">人工智能</span></a><a href="https://m.weibo.cn/p/index?extparam=%E7%A1%95%E5%A3%AB%E8%AE%BA%E6%96%87&amp;containerid=1008084cacf38f5903dc7b04550404d0bd3608" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://n.sinaimg.cn/photo/5213b46e/20180926/timeline_card_small_super_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">硕士论文</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%BC%80%E6%BA%90%23" data-hide=""><span class="surl-text">#开源#</span></a><img style="" src="https://tvax4.sinaimg.cn/large/6f830abfly1hwgrb57m70j22cl18bnpd.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

本文介绍了一种名为LiFT的新微调方法，旨在通过利用人工反馈来优化文本到视频生成模型的匹配度。研究团队首先构建了一个包含1万个人工评分及其理由的数据集LiFT-HRA，并基于此训练了一个奖励模型LiFT-Critic，该模型能够学习奖励函数，作为人类判断的代理。最后，利用学到的奖励函数通过最大化奖励加权的似然性来调整T2V模型。实验结果显示，经过微调的模型在所有16项指标上均优于CogVideoX-5B，证明了人类反馈在提高生成视频匹配度和质量方面的潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-20T03:17:56Z
- **目录日期**: 2025-03-20
