# #港中文教AI用视觉工具辅助思考##提升AI看图推理能力新方法#教AI学会使用工具，看图推理就能变得更强？！港中文、微软等机构推出了OpenThinkIMG，让AI像人一样“...

**URL**: https://weibo.com/6105753431/PseuR3pLZ

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%B8%AF%E4%B8%AD%E6%96%87%E6%95%99AI%E7%94%A8%E8%A7%86%E8%A7%89%E5%B7%A5%E5%85%B7%E8%BE%85%E5%8A%A9%E6%80%9D%E8%80%83%23&amp;extparam=%23%E6%B8%AF%E4%B8%AD%E6%96%87%E6%95%99AI%E7%94%A8%E8%A7%86%E8%A7%89%E5%B7%A5%E5%85%B7%E8%BE%85%E5%8A%A9%E6%80%9D%E8%80%83%23" data-hide=""><span class="surl-text">#港中文教AI用视觉工具辅助思考#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E6%8F%90%E5%8D%87AI%E7%9C%8B%E5%9B%BE%E6%8E%A8%E7%90%86%E8%83%BD%E5%8A%9B%E6%96%B0%E6%96%B9%E6%B3%95%23&amp;extparam=%23%E6%8F%90%E5%8D%87AI%E7%9C%8B%E5%9B%BE%E6%8E%A8%E7%90%86%E8%83%BD%E5%8A%9B%E6%96%B0%E6%96%B9%E6%B3%95%23" data-hide=""><span class="surl-text">#提升AI看图推理能力新方法#</span></a><br><br>教AI学会使用工具，看图推理就能变得更强？！<br><br>港中文、微软等机构推出了OpenThinkIMG，让AI像人一样“动手操作”，学会用视觉工具来辅助思考。<br><br>用这个框架训练出的AI，在图表推理任务上不仅大幅超越了传统训练方法，甚至比GPT-4.1还强，和Gemini表现打平，而且参数更小。【图2】<br><br>三大核心功能：<br><br>1. 模块化工具部署：不管是GroundingDINO、SAM这种热门工具，还是你自家开发的，都能接入OpenThinkIMG。还能独立运行、按需调用。<br>    <br>2. 智能体训练机制：融合了新算法V-ToolRL，让AI能通过“试错+奖励”自己摸索怎么用工具。训练过程是先监督微调（SFT），再强化实战。<br>    <br>3. 高质量数据生成：GPT-4o先做任务规划，OpenThinkIMG执行并记录，然后AI+人工联合质检，确保训练数据高质量。<br><br>在饼图分析任务中，V-ToolRL会自动放大关键区域、用OCR识别数值，准确计算差异；面对折线图，还能调用多种工具辅助对比找趋势。【图3】<br><br>未来，团队将扩展更多工具和任务场景，推动AI更深入“图像思考”方向发展。<br><br>技术报告：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fpdf%2F2505.08617" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>GitHub仓库：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fzhaochen0110%2FOpenThinkIMG" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>数据集和模型：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fhuggingface.co%2Fcollections%2FWarrieryes%2Fopenthinkimg-68244a63e97a24d9b7ffcde9" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1ify1kwodj30zk0iqne2.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3ly1i1ify6abzdj30zk0cvgqo.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3ly1i1ifyapqh5j30zk0q84qp.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

香港中文大学与微软等机构合作开发了OpenThinkIMG框架，通过视觉工具辅助AI进行图像推理。该框架采用模块化工具部署、智能体训练机制（V-ToolRL算法）和高质量数据生成三大核心技术，使AI能像人类一样调用工具分析图表。实验显示，其性能超越传统方法及GPT-4.1，与Gemini相当且参数更小。例如在饼图分析中能自动放大关键区域并OCR识别数值，折线图分析中可多工具协同找趋势。项目已开源技术报告、代码及模型，未来将扩展更多工具应用场景。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-17T06:02:51Z
- **目录日期**: 2025-05-17
