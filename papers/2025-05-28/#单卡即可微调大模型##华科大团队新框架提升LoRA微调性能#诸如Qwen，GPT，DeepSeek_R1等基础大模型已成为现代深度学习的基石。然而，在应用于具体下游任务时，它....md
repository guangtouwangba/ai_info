# #单卡即可微调大模型##华科大团队新框架提升LoRA微调性能#诸如Qwen，GPT，DeepSeek R1等基础大模型已成为现代深度学习的基石。然而，在应用于具体下游任务时，它...

**URL**: https://weibo.com/6105753431/PtTAf6qLJ

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%95%E5%8D%A1%E5%8D%B3%E5%8F%AF%E5%BE%AE%E8%B0%83%E5%A4%A7%E6%A8%A1%E5%9E%8B%23&amp;extparam=%23%E5%8D%95%E5%8D%A1%E5%8D%B3%E5%8F%AF%E5%BE%AE%E8%B0%83%E5%A4%A7%E6%A8%A1%E5%9E%8B%23" data-hide=""><span class="surl-text">#单卡即可微调大模型#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%8D%8E%E7%A7%91%E5%A4%A7%E5%9B%A2%E9%98%9F%E6%96%B0%E6%A1%86%E6%9E%B6%E6%8F%90%E5%8D%87LoRA%E5%BE%AE%E8%B0%83%E6%80%A7%E8%83%BD%23&amp;extparam=%23%E5%8D%8E%E7%A7%91%E5%A4%A7%E5%9B%A2%E9%98%9F%E6%96%B0%E6%A1%86%E6%9E%B6%E6%8F%90%E5%8D%87LoRA%E5%BE%AE%E8%B0%83%E6%80%A7%E8%83%BD%23" data-hide=""><span class="surl-text">#华科大团队新框架提升LoRA微调性能#</span></a><br><br>诸如Qwen，GPT，DeepSeek R1等基础大模型已成为现代深度学习的基石。<br><br>然而，在应用于具体下游任务时，它们庞大的参数规模使得额外微调成本较高。<br><br>为了解决这一问题，近期的研究聚焦于低秩适应 (LoRA) 方法，通过保持基座模型参数冻结，仅对新增的小型轻量级适配器进行微调，从而降低微调成本。<br><br>尽管LoRA具有较高的效率，然而其微调性能往往不及全量微调。<br><br>面对这一挑战，华中科技大学和香港中文大学团队提出了一项全新的LoRA微调框架——GOAT，该工作已成功被ICML 2025正式接收。<br><br>这项研究提出了一套自适应奇异值初始化与混合专家梯度对齐策略，成功缓解低秩适应（LoRA）性能不足的难题，在25个多领域任务中实现接近甚至超越全参数微调（Full FT）的效果，同时仅需调整极小比例参数。<br><br>更多细节：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fmp.weixin.qq.com%2Fs%2F2ZvHV0YBTWUkXJFkmJ2D6g" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">单卡即可微调大模型！内存占用仅1/8，性能依然拉满 | ICML 2025</span></a><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i1v3624cv7j30jg046tac.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1v36630spj30jg0aojw4.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

华中科技大学与香港中文大学团队提出新型LoRA微调框架GOAT，被ICML 2025接收。该框架通过自适应奇异值初始化和混合专家梯度对齐策略，显著提升低秩适应（LoRA）性能，在25个多领域任务中接近或超越全参数微调效果，同时仅需调整极少量参数。实验显示单张显卡即可高效微调大模型，内存占用仅为传统方法的1/8，解决了大模型下游任务微调成本高的问题。该方法为Qwen、GPT等大模型的轻量化微调提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T05:03:05Z
- **目录日期**: 2025-05-28
