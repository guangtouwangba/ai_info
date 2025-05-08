# #多模态奖励瓶颈再突破##奖励模型也能学推理#奖励模型能做得多好？中科院、清华、快手和南大联合发布的R1-Reward模型，借助强化学习技巧，把多模态奖励模型的表...

**URL**: https://weibo.com/6105753431/PqT2uEc3I

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A4%9A%E6%A8%A1%E6%80%81%E5%A5%96%E5%8A%B1%E7%93%B6%E9%A2%88%E5%86%8D%E7%AA%81%E7%A0%B4%23&amp;extparam=%23%E5%A4%9A%E6%A8%A1%E6%80%81%E5%A5%96%E5%8A%B1%E7%93%B6%E9%A2%88%E5%86%8D%E7%AA%81%E7%A0%B4%23" data-hide=""><span class="surl-text">#多模态奖励瓶颈再突破#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%A5%96%E5%8A%B1%E6%A8%A1%E5%9E%8B%E4%B9%9F%E8%83%BD%E5%AD%A6%E6%8E%A8%E7%90%86%23&amp;extparam=%23%E5%A5%96%E5%8A%B1%E6%A8%A1%E5%9E%8B%E4%B9%9F%E8%83%BD%E5%AD%A6%E6%8E%A8%E7%90%86%23" data-hide=""><span class="surl-text">#奖励模型也能学推理#</span></a><br><br>奖励模型能做得多好？<br><br>中科院、清华、快手和南大联合发布的R1-Reward模型，借助强化学习技巧，把多模态奖励模型的表现提升了5%-15%。推理时多做几次，效果还能再往上涨！<br><br>这波升级背后，是他们把训练奖励模型这事儿，重新定义为一个规则驱动的强化学习问题：给模型一个问题和两个答案，让它学会判断并说出理由。<br><br>他们设计的新算法叫StableReinforce，有三大关键：<br><br>- 预裁剪（Pre-Clip）：防止训练过程爆炸；<br><br>- 优势过滤器（Advantage Filter）：防止极端样本扰乱训练；<br><br>- 一致性奖励：引入“裁判”模型检查模型分析和结论是否一致，鼓励逻辑自洽。<br><br>训练流程也很讲究——先用GPT-4o生成20万条样本的思考过程做“预习”，再用GPT-4o判断哪些样本难度大，在强化学习阶段重点攻坚这些“硬骨头”。<br><br>结果确实不负众望：在VLReward Bench等多个榜单上全面超越SOTA，推理时做多次采样还能进一步拉高准确率。甚至在“只要有一次答对就算对”的Any Correct评估方式下，R1-Reward准确率几乎能接近100%。<br><br>更神奇的是，这个模型还学会了反思和纠错，推理过程还变得更短、更高效了。详情请见文章： <a href="https://weibo.com/ttarticle/p/show?id=2309405164081026826368" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">突破多模态奖励瓶颈！中科院清华快手用强化学习赋予模型长期推理能力</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i187jztqccj30n00cydhh.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

中科院、清华、快手和南大联合研发的R1-Reward模型通过强化学习技术，将多模态奖励模型性能提升5%-15%。该研究将奖励模型训练重构为规则驱动的强化学习问题，提出StableReinforce算法，包含预裁剪、优势过滤器和一致性奖励三大关键技术。模型先通过GPT-4o生成的20万样本进行预训练，再针对高难度样本重点优化。实验显示，该模型在多个基准测试中超越现有最优方法，推理效率更高，在Any Correct评估下准确率接近100%，并展现出反思纠错能力。论文详见链接。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T13:12:53Z
- **目录日期**: 2025-05-08
