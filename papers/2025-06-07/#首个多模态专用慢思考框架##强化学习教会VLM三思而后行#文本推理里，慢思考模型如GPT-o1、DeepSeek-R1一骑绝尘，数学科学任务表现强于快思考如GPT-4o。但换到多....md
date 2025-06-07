# #首个多模态专用慢思考框架##强化学习教会VLM三思而后行#文本推理里，慢思考模型如GPT-o1、DeepSeek-R1一骑绝尘，数学科学任务表现强于快思考如GPT-4o。但换到多...

**URL**: https://weibo.com/6105753431/PvrVLBLGR

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E9%A6%96%E4%B8%AA%E5%A4%9A%E6%A8%A1%E6%80%81%E4%B8%93%E7%94%A8%E6%85%A2%E6%80%9D%E8%80%83%E6%A1%86%E6%9E%B6%23&amp;extparam=%23%E9%A6%96%E4%B8%AA%E5%A4%9A%E6%A8%A1%E6%80%81%E4%B8%93%E7%94%A8%E6%85%A2%E6%80%9D%E8%80%83%E6%A1%86%E6%9E%B6%23" data-hide=""><span class="surl-text">#首个多模态专用慢思考框架#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0%E6%95%99%E4%BC%9AVLM%E4%B8%89%E6%80%9D%E8%80%8C%E5%90%8E%E8%A1%8C%23&amp;extparam=%23%E5%BC%BA%E5%8C%96%E5%AD%A6%E4%B9%A0%E6%95%99%E4%BC%9AVLM%E4%B8%89%E6%80%9D%E8%80%8C%E5%90%8E%E8%A1%8C%23" data-hide=""><span class="surl-text">#强化学习教会VLM三思而后行#</span></a><br><br>文本推理里，慢思考模型如GPT-o1、DeepSeek-R1一骑绝尘，数学科学任务表现强于快思考如GPT-4o。但换到多模态推理，这些“反思大师”却落了下风，不仅没赢，甚至被Qwen2.5-VL-72B反超。<br><br>为啥慢思考在图文场景里不灵了？港科大等团队研究发现，视觉语言模型（VLM）构建慢思考能力卡在两个点：优势消失和反思惰性。<br><br>为破解难题，团队整理出39K条高质量多模态推理数据（ViRL39K），并提出“VL-Rethinker”框架，靠两项关键机制提升能力：<br><br>1. 优势样本回放（SSR）：动态存储并优先复用非零优势的关键样本，比如“难题做对”或“简单题做错”。这样训练聚焦高价值样本，提升稳定性和效率。<br><br>2. 强制反思机制：每次回答后强制加一句“反思提示”，引导模型自我验证、纠错、提问。不是每题都反思，而是学会何时该反思。结果发现，模型连题目本身的错误也能识别出来了。<br><br>VL-Rethinker在MathVista、MathVerse、MMMU-Pro等数据集全面超越GPT-o1，表现接近OpenAI-o1，领先现有开源模型。不同规模模型（72B与7B）均大幅提升，验证了“慢思考”在多模态领域的潜力。 <a href="https://weibo.com/ttarticle/p/show?id=2309405174955376771517" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">首个多模态专用慢思考框架！强化学习教会VLM「三思而后行」</span></a><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3ly1i26wixx0wwj30rs0fmwfx.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

港科大等团队针对视觉语言模型(VLM)在多模态推理中的"慢思考"短板，提出创新框架VL-Rethinker。该方案通过两项机制突破传统局限：1）优势样本回放(SSR)精选高价值训练样本；2）强制反思机制引导模型自主判断反思时机。基于39K条多模态数据(ViRL39K)的实验显示，该框架在MathVista等基准测试中超越GPT-o1，接近OpenAI-o1水平，不同规模模型均获显著提升。研究证实"慢思考"策略经优化后同样适用于多模态领域，使VLM具备自主纠错和问题识别能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-07T16:03:27Z
- **目录日期**: 2025-06-07
