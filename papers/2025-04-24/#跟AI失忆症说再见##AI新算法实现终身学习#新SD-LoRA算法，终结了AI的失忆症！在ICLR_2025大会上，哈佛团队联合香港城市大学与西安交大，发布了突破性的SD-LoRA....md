# #跟AI失忆症说再见##AI新算法实现终身学习#新SD-LoRA算法，终结了AI的失忆症！在ICLR 2025大会上，哈佛团队联合香港城市大学与西安交大，发布了突破性的SD-LoRA...

**URL**: https://weibo.com/6105753431/PoJ1yuUDQ

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E8%B7%9FAI%E5%A4%B1%E5%BF%86%E7%97%87%E8%AF%B4%E5%86%8D%E8%A7%81%23&amp;extparam=%23%E8%B7%9FAI%E5%A4%B1%E5%BF%86%E7%97%87%E8%AF%B4%E5%86%8D%E8%A7%81%23" data-hide=""><span class="surl-text">#跟AI失忆症说再见#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23AI%E6%96%B0%E7%AE%97%E6%B3%95%E5%AE%9E%E7%8E%B0%E7%BB%88%E8%BA%AB%E5%AD%A6%E4%B9%A0%23&amp;extparam=%23AI%E6%96%B0%E7%AE%97%E6%B3%95%E5%AE%9E%E7%8E%B0%E7%BB%88%E8%BA%AB%E5%AD%A6%E4%B9%A0%23" data-hide=""><span class="surl-text">#AI新算法实现终身学习#</span></a><br><br>新SD-LoRA算法，终结了AI的失忆症！<br><br>在ICLR 2025大会上，哈佛团队联合香港城市大学与西安交大，发布了突破性的SD-LoRA算法，是一种rehearsal-free、推理高效、可端到端优化的持续学习方法。<br><br>该技术通过固定已学习任务的方向参数，仅调整幅度权重，避免了历史数据的存储需求，同时减少了50%以上的参数存储量。<br><br>具体来说，作者提出了⼀种⾯向预训练⼤模型的持续学习⽅法SD-LoRA，具备⽆需回放（rehearsal-free）、推理⾼效、可端到端训练等优点，并进⼀步设计了两个参数更⾼效的变体。<br><br>本文从理论与实证层⾯深入分析了SD-LoRA的⼯作机制，解释其如何避免依赖任务特定模块的选择（Prompt或者LoRA） ，为持续学习提供了新的解决思路和⽅案。<br><br>在多个基准任务和预训练模型上的实验表明，SD-LoRA在不增加推理开销的前提下显著缓解了灾难性遗忘问题。<br><br>标志着AI模型持续学习的一大进步。<br><br>与传统的混合专家模型不同，SD-LoRA采用低秩矩阵分解，使得学习过程更为高效且无需回放历史任务数据，显著减轻了灾难性遗忘问题。而且，SD-LoRA无需依赖复杂的路由机制，优化了推理效率，并能端到端训练。<br><br>实验结果表明，在多个基准测试和预训练模型上，SD-LoRA不仅保持了高准确率，还有效减少了推理开销。详情请见文章： <a href="https://weibo.com/ttarticle/p/show?id=2309405158929490641116" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">告别“AI失忆症”！新型SD-LoRA算法实现终身学习｜ICLR 2025</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0rqxkuo5tj30fe08ojrm.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

哈佛团队联合香港城市大学与西安交大在ICLR 2025大会上发布了SD-LoRA算法，解决了AI持续学习中的"失忆症"问题。该算法通过固定已学任务方向参数、仅调整幅度权重，无需存储历史数据，减少50%以上参数存储量。SD-LoRA具有无回放、高效推理和端到端训练优势，采用低秩矩阵分解替代复杂路由机制，显著缓解灾难性遗忘问题。实验表明，该算法在多个基准测试中保持高准确率的同时降低了推理开销，为AI终身学习提供了新方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T08:29:52Z
- **目录日期**: 2025-04-24
