# #用LLM做CAD##用自然语言建模CAD#当你用CAD做3D模型时，有没有想过用AI自动画图？程序员Will Patrick不仅想到，还真正做到了。这不仅是一次技术实验，更是一次跨...

**URL**: https://weibo.com/6105753431/PoKQo0XOW

## 原始摘要

<a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%A8LLM%E5%81%9ACAD%23&amp;extparam=%23%E7%94%A8LLM%E5%81%9ACAD%23" data-hide=""><span class="surl-text">#用LLM做CAD#</span></a><a href="https://m.weibo.cn/search?containerid=231522type%3D1%26t%3D10%26q%3D%23%E7%94%A8%E8%87%AA%E7%84%B6%E8%AF%AD%E8%A8%80%E5%BB%BA%E6%A8%A1CAD%23&amp;extparam=%23%E7%94%A8%E8%87%AA%E7%84%B6%E8%AF%AD%E8%A8%80%E5%BB%BA%E6%A8%A1CAD%23" data-hide=""><span class="surl-text">#用自然语言建模CAD#</span></a><br><br>当你用CAD做3D模型时，有没有想过用AI自动画图？<br><br>程序员Will Patrick不仅想到，还真正做到了。这不仅是一次技术实验，更是一次跨领域探索。<br><br>Will Patrick一开始并没有直接从传统的CAD工具入手，而是选择了OpenSCAD。<br><br>与Fusion 360或Solidworks等工具不同，OpenSCAD是一个程序化的CAD工具，靠代码就能定义3D模型的各个部分。<br><br>他因此产生了一个想法：如果能够让大语言模型（LLM）自动生成OpenSCAD代码，那么机械制图就能实现自动化。<br><br>Will Patrick开始了他的实验：让LLM学习如何生成OpenSCAD代码来创建3D模型。<br><br>他启动AI编程工具Cursor，首先向它提出了一个简单的请求：“创建一个iPhone外壳”。<br><br>AI还真的生成了一个iPhone模型，跑出来的效果be like：【图1】<br><br>随着实验的深入，Will意识到，单纯让AI生成模型并不足够，AI还需要理解基本的机械设计原理。<br><br>因此，他让LLM学习一些常见的设计规范和工程原理。<br><br>例如，AI生成一个带螺栓孔的法兰（flange，一种器件【图2】）时，AI还自动计算了螺栓孔的大小、数量和排列方式，很有工程制图那味儿。【图3】<br><br>为了确保这些模型符合工程要求，Will还设计了一个自动化评估系统。【图4】<br><br>这个系统通过将AI生成的模型与参考模型进行对比，自动检查几何精度，确保模型符合体积、边界框等一系列几何标准。<br><br>通过这次实验，Will不仅完成了技术突破，更为CAD建模的未来发展提供了全新的视角。<br><br>他认为，随着技术不断进步，未来的设计工具将不再局限于传统的手动操作，而是通过GenCAD（生成式CAD）实现更加智能化的设计流程。<br><br>届时，设计师只需告诉AI：“在这个位置需要一个支架”，AI便能自动生成与现有组件完全匹配的支架模型。<br><br>感兴趣的小伙伴可以点击原文，有具体的执行步骤：<a href="https://weibo.cn/sinaurl?u=https%3A%2F%2Fwillpatrick.xyz%2Ftechnology%2F2025%2F04%2F23%2Fteaching-llms-how-to-solid-model.html" data-hide=""><span class="url-icon"><img style="width: 1rem;height: 1rem" src="https://h5.sinaimg.cn/upload/2015/09/25/3/timeline_card_small_web_default.png" referrerpolicy="no-referrer"></span><span class="surl-text">网页链接</span></a><img style="" src="https://tvax3.sinaimg.cn/large/006Fd7o3gy1i0s006tptmg30hs0biu0b.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax1.sinaimg.cn/large/006Fd7o3gy1i0s00c3q1dj30qp0zke81.jpg" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax4.sinaimg.cn/large/006Fd7o3gy1i0s00899l6g30hs0akqfz.gif" referrerpolicy="no-referrer"><br><br><img style="" src="https://tvax2.sinaimg.cn/large/006Fd7o3gy1i0s00hot8lj31660oygth.jpg" referrerpolicy="no-referrer"><br><br>

## AI 摘要

程序员Will Patrick探索了用大语言模型(LLM)生成OpenSCAD代码来自动创建3D模型的方法。通过AI编程工具Cursor，成功实现了从自然语言指令(如"创建iPhone外壳")生成3D模型，并进一步让AI学习机械设计原理，自动计算螺栓孔等工程参数。他还开发了自动化评估系统来验证模型的几何精度。这项实验展示了"生成式CAD"(GenCAD)的潜力，未来设计师可能只需描述需求，AI就能自动生成匹配的3D模型，为CAD设计带来智能化革新。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-24T12:04:08Z
- **目录日期**: 2025-04-24
