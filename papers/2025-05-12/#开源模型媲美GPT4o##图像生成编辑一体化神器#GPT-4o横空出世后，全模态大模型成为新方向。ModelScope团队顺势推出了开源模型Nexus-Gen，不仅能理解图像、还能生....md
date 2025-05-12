# #开源模型媲美GPT4o##图像生成编辑一体化神器#GPT-4o横空出世后，全模态大模型成为新方向。ModelScope团队顺势推出了开源模型Nexus-Gen，不仅能理解图像、还能生...

**URL**: https://weibo.com/6105753431/Prno8yySk

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%BC%80%E6%BA%90%E6%A8%A1%E5%9E%8B%E5%AA%B2%E7%BE%8EGPT4o%23&amp;extparam=%23%E5%BC%80%E6%BA%90%E6%A8%A1%E5%9E%8B%E5%AA%B2%E7%BE%8EGPT4o%23" data-hide=""><span class="surl-text">#开源模型媲美GPT4o#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E5%9B%BE%E5%83%8F%E7%94%9F%E6%88%90%E7%BC%96%E8%BE%91%E4%B8%80%E4%BD%93%E5%8C%96%E7%A5%9E%E5%99%A8%23&amp;extparam=%23%E5%9B%BE%E5%83%8F%E7%94%9F%E6%88%90%E7%BC%96%E8%BE%91%E4%B8%80%E4%BD%93%E5%8C%96%E7%A5%9E%E5%99%A8%23" data-hide=""><span class="surl-text">#图像生成编辑一体化神器#</span></a><br><br>GPT-4o横空出世后，全模态大模型成为新方向。ModelScope团队顺势推出了开源模型Nexus-Gen，不仅能理解图像、还能生成和编辑，且图像质量直逼GPT-4o。<br><br>Nexus-Gen采用Transformer加扩散模型的技术路径，融合MLLM的语言理解能力和扩散模型的图像渲染能力。与传统All-to-All模型不同，它在高维特征空间进行建模，显著提升图像质量。<br><br>为了解决图像生成中误差积累的问题，团队提出“预填充自回归”策略，优化训练与推理一致性，从而提升预测准确率。<br><br>训练方面，Nexus-Gen统一了三类任务的数据格式，并使用了25M规模的数据，包括图像理解、生成和编辑任务。图像编辑部分尤其依赖ImagePulse数据集，涵盖添加、去除、风格迁移等操作。自回归模块和扩散模块分开训练，分别使用SWIFT和DiffSynth-Studio框架。<br><br>模型目前在图像理解、生成和编辑三方面均展示了出色能力，支持多prompt、故事线式编辑等操作，具有广阔的应用潜力。未来团队还计划在模型规模、图像Token数量、融合训练等方向持续优化。<br><br>目前包括论文、代码、模型和数据集都已开源——<br><br>论文链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Farxiv.org%2Fpdf%2F2504.21356" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a>  <br>代码链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fgithub.com%2Fmodelscope%2FNexus-Gen" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a>  <br>模型链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.modelscope.cn%2Fmodels%2FDiffSynth-Studio%2FNexus-Gen" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a>  <br>数据集（ImagePulse）链接：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwww.modelscope.cn%2Fcollections%2FImagePulse----tulvmaidong-7c3b8283a43e40" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a> <a href="https://weibo.com/ttarticle/p/show?id=2309405165247647580712" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_article_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">开源全能图像模型媲美GPT-4o！理解生成编辑同时搞定，解决扩散模型误</span></a><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i1bxlr3gb6j30rs0fm0wz.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

ModelScope团队推出开源多模态模型Nexus-Gen，性能媲美GPT-4o。该模型结合Transformer和扩散模型技术，在高维特征空间进行建模，显著提升图像质量。采用"预填充自回归"策略解决误差积累问题，统一了图像理解、生成和编辑三大任务。训练使用25M规模数据，其中编辑任务依赖ImagePulse数据集。模型支持多prompt处理和故事线编辑，目前论文、代码、模型及数据集均已开源。未来计划在模型规模等方面继续优化。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-12T01:30:25Z
- **目录日期**: 2025-05-12
