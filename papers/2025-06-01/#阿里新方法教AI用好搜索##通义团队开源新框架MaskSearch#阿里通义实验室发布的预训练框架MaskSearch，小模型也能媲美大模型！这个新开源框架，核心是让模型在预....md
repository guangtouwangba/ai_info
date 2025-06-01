# #阿里新方法教AI用好搜索##通义团队开源新框架MaskSearch#阿里通义实验室发布的预训练框架MaskSearch，小模型也能媲美大模型！这个新开源框架，核心是让模型在预...

**URL**: https://weibo.com/6105753431/PumYhswig

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%98%BF%E9%87%8C%E6%96%B0%E6%96%B9%E6%B3%95%E6%95%99AI%E7%94%A8%E5%A5%BD%E6%90%9C%E7%B4%A2%23&amp;extparam=%23%E9%98%BF%E9%87%8C%E6%96%B0%E6%96%B9%E6%B3%95%E6%95%99AI%E7%94%A8%E5%A5%BD%E6%90%9C%E7%B4%A2%23" data-hide=""><span class="surl-text">#阿里新方法教AI用好搜索#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%80%9A%E4%B9%89%E5%9B%A2%E9%98%9F%E5%BC%80%E6%BA%90%E6%96%B0%E6%A1%86%E6%9E%B6MaskSearch%23&amp;extparam=%23%E9%80%9A%E4%B9%89%E5%9B%A2%E9%98%9F%E5%BC%80%E6%BA%90%E6%96%B0%E6%A1%86%E6%9E%B6MaskSearch%23" data-hide=""><span class="surl-text">#通义团队开源新框架MaskSearch#</span></a><br><br>阿里通义实验室发布的预训练框架MaskSearch，小模型也能媲美大模型！<br><br>这个新开源框架，核心是让模型在预训练阶段练习“边推理边用搜索引擎”。方法来源于BERT的“掩码预测”思路，但MaskSearch不是让模型凭空猜词，而是逼它学会通过搜索获取外部知识来补全被遮蔽的信息。这类任务叫做RAMP（检索增强型掩码预测）。<br><br>MaskSearch特别之处在于，它能同时兼容监督微调（SFT）和强化学习（RL）训练方式。在SFT中，它用“多智能体+教师模型”的方式自动生成高质量思维链；在RL中，它靠“混合奖励”系统引导模型优化推理路径，避免只为得分堆信息。<br><br>更进一步，作者还设计了课程学习策略，让模型从简单任务逐渐适应复杂场景，显著提升模型泛化能力。<br><br>实验结果亮眼。在多个问答数据集上，MaskSearch都比传统训练方式效果更好。尤其是在跨领域任务上，小模型的表现甚至能媲美大模型。<br><br>同时，研究也发现不同掩码策略、奖励函数对最终性能有明显影响，说明“让模型学会怎么学”本身也是个技术活。 <a href="https://weibo.com/ttarticle/p/show?id=2309405172381458235537" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">阿里通义开源「推理+搜索」预训练新框架：小模型媲美大模型</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1yowvo0xgj30rs0fm0x3.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

阿里通义实验室开源了新型预训练框架MaskSearch，通过检索增强型掩码预测(RAMP)技术，让AI模型在训练时学会结合搜索引擎进行推理。该框架创新性地融合了监督微调(SFT)和强化学习(RL)两种训练方式，采用多智能体协作生成思维链，并设计混合奖励机制优化推理路径。实验显示，采用MaskSearch的小模型在跨领域问答任务中表现优异，甚至能媲美大模型性能。研究还发现掩码策略和奖励函数对模型效果有重要影响，表明训练方法本身需要精细设计。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-01T00:03:14Z
- **目录日期**: 2025-06-01
