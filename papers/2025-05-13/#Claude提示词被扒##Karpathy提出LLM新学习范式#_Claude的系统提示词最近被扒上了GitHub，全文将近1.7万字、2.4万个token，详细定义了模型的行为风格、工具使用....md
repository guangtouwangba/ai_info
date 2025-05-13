# #Claude提示词被扒##Karpathy提出LLM新学习范式# Claude的系统提示词最近被扒上了GitHub，全文将近1.7万字、2.4万个token，详细定义了模型的行为风格、工具使用...

**URL**: https://weibo.com/6105753431/PrClwitVm

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Claude%E6%8F%90%E7%A4%BA%E8%AF%8D%E8%A2%AB%E6%89%92%23&amp;extparam=%23Claude%E6%8F%90%E7%A4%BA%E8%AF%8D%E8%A2%AB%E6%89%92%23" data-hide=""><span class="surl-text">#Claude提示词被扒#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23Karpathy%E6%8F%90%E5%87%BALLM%E6%96%B0%E5%AD%A6%E4%B9%A0%E8%8C%83%E5%BC%8F%23&amp;extparam=%23Karpathy%E6%8F%90%E5%87%BALLM%E6%96%B0%E5%AD%A6%E4%B9%A0%E8%8C%83%E5%BC%8F%23" data-hide=""><span class="surl-text">#Karpathy提出LLM新学习范式#</span></a> <br><br>Claude的系统提示词最近被扒上了GitHub，全文将近1.7万字、2.4万个token，详细定义了模型的行为风格、工具使用方式、回答策略等等。【图1】<br><br>虽然被称作“泄漏”，但Claude的提示词本身在官方页面上是可查的。【图2】<br><br>而真正引发热议的，是大神Karpathy的点评。他认为这份Claude提示词文档，暴露了LLM目前在“学习方式”上的核心问题。<br><br>他说，当前主流的大模型学习方式，还停留在“像RLHF（强化学习）这种方法调完就完”的阶段，缺乏模拟人类经验积累的机制。于是他提出一个新思路：系统提示学习（System Prompt Learning）。【图3】<br><br>简单来说，新范式的思路就是——<br><br>- 用系统提示给LLM写“备忘录”，帮助它记住问题解决策略；<br><br>- 像人一样形成长期记忆，遇到相似问题能快速调取并泛化；<br><br>- 提示词不仅是“设定”，还能成为模型不断调整自我行为的“自主策略层”。<br><br>这种方式的好处在于，提示词可以随时人工检查、更新，既具备“测试时训练”的能力，又比权重微调更安全灵活。<br><br>然而反对者认为，这种方法只是变相“手动喂知识”，大模型自身并不理解提示内容，是否真的形成“学习”还有待验证。【图4】<br><br>回到Claude这份超长提示词来，里面大约80%的内容都和工具使用有关，比如如何引用、如何使用搜索、遇到日期类问题该怎么处理等等。还有一些是热修内容，补充知识截止点之后的重要信息。<br><br>这种“全手工打造”的提示体系，让Claude有了不一样的体验。<br><br>正因如此，更凸显出LLM在自我学习、自主演化上的局限。而Karpathy提出的新思路，或许会是下一波AI范式的起点。<br><br>Claude提示词泄漏版：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fraw.githubusercontent.com%2Fasgeirtj%2Fsystem_prompts_leaks%2Frefs%2Fheads%2Fmain%2Fclaude.txt" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><br>Claude提示词官方版：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fdocs.anthropic.com%2Fen%2Frelease-notes%2Fsystem-prompts" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1drn4s8cbj30wg0esjwl.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1drn6f7alj321m1dah4a.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i1drn7vrdjj314u18e1ht.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i1drn8yr7kj30wk0uwqic.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

Claude的1.7万字系统提示词被公开，其中80%内容涉及工具使用规范。AI专家Karpathy指出这暴露了当前大模型依赖静态RLHF训练的局限，提出"系统提示学习"新范式——通过可更新的提示词让模型形成类人记忆机制，实现持续学习。支持者认为这种方法比微调更安全可控，反对者则质疑其是否真正实现"理解"。该事件突显了LLM在自主演化方面的不足，可能推动AI学习方式的革新。Claude官方表示提示词本就是公开文档，并非泄漏。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T08:04:28Z
- **目录日期**: 2025-05-13
